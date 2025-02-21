package db

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB
var once sync.Once

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return instance
}

// InitDB 初始化数据库连接
func InitDB() error {
	var initDBErr error
	once.Do(func() {
		// 获取用户主目录
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return
		}

		// 创建应用数据目录
		dataDir := filepath.Join(homeDir, ".todo_list")
		if err = os.MkdirAll(dataDir, 0755); err != nil {
			return
		}

		// 数据库文件路径
		dbPath := filepath.Join(dataDir, "todo.db")

		// 配置GORM日志
		logConfig := logger.Config{
			SlowThreshold:             0,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}

		// 连接数据库
		db, dbErr := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logConfig),
		})
		if dbErr != nil {
			initDBErr = dbErr
			return
		}

		// 自动迁移数据库结构
		if err = db.AutoMigrate(&Todo{}, &TodoGroup{}); err != nil {
			return
		}

		instance = db
	})
	return initDBErr
}

// CloseDB
func CloseDB() error {
	sqlDB, err := GetDB().DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
