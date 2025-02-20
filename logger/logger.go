package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var logger *slog.Logger

// InitLogger 初始化日志记录器
func InitLogger() error {
	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("获取用户主目录失败: %w", err)
	}

	// 创建日志目录
	dataDir := filepath.Join(homeDir, ".todo_list", "logs")
	if err = os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("创建日志目录失败: %w", err)
	}

	// 创建日志文件
	logFileName := fmt.Sprintf("todo_list_%s.log", time.Now().Format("2006-01-02"))
	logFilePath := filepath.Join(dataDir, logFileName)
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("创建日志文件失败: %w", err)
	}

	// 配置slog
	handler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})

	logger = slog.New(handler)
	slog.SetDefault(logger)

	return nil
}
