package db

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Todo 待办事项模型
type Todo struct {
	Id          uint      `gorm:"column:id;primaryKey;autoIncrement;comment:编码;" json:"id"`      // 编码
	Title       string    `gorm:"column:title;not null;comment:标题;" json:"title"`                // 标题
	Description string    `gorm:"column:description;default:'';comment:描述;" json:"description"`  // 描述
	Completed   bool      `gorm:"column:completed;default:false;comment:完成状态;" json:"completed"` // 完成状态
	DueDate     time.Time `gorm:"column:due_date;comment:开始时间;" json:"due_date"`                 // 添加截止日期字段
	CreatedAt   time.Time `gorm:"column:created_at;comment:创建时间;" json:"created_at"`             // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;comment:更新时间;" json:"updated_at"`             // 更新时间
}

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
		if err = db.AutoMigrate(&Todo{}); err != nil {
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

// CreateTodo 创建待办事项
func CreateTodo(todo *Todo) error {
	return GetDB().Create(todo).Error
}

// GetTodos 获取所有待办事项
func GetTodos() ([]Todo, error) {
	var todos []Todo
	err := GetDB().Order("due_date asc").Find(&todos).Error
	return todos, err
}

// GetTodosByDate 获取指定日期的待办事项
func GetTodosByDate(date time.Time) ([]Todo, error) {
	var todos []Todo
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := GetDB().Where("due_date BETWEEN ? AND ?", startOfDay, endOfDay).Order("due_date asc").Find(&todos).Error
	return todos, err
}

// UpdateTodo 更新待办事项
func UpdateTodo(todo *Todo) error {
	return GetDB().Where("id = ?", todo.Id).Save(todo).Error
}

// DeleteTodo 删除待办事项
func DeleteTodo(id uint) error {
	return GetDB().Delete(&Todo{}, id).Error
}
