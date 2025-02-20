package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"todo_list/db"
	"todo_list/logger"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// 初始化日志系统
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("初始化日志系统失败: %v\n", err)
		return
	}
	slog.Info("应用程序启动成功")
}

// shutdown is called when the app is closing
func (a *App) shutdown(ctx context.Context) {
	if err := db.CloseDB(); err != nil {
		slog.Error("关闭数据库失败", "错误信息", err.Error())
	}
	slog.Info("应用程序关闭")
}

// CreateTodo 创建新的待办事项
func (a *App) CreateTodo(title string, level string, dueDate string) error {
	slog.Info("创建待办事项", "title", title, "dueDate", dueDate)

	todo := &db.Todo{
		Title: title,
		Level: level,
	}

	if dueDate != "" {
		parsedTime, err := time.Parse("2006-01-02T15:04", dueDate)
		if err != nil {
			slog.Error("日期格式无效", "错误信息", err.Error())
			return fmt.Errorf("invalid date format: %v", err)
		}

		todo.DueDate = parsedTime
	} else {
		todo.DueDate = time.Now()
	}

	return db.CreateTodo(todo)
}

// UpdateTodo 更新待办事项
func (a *App) UpdateTodo(id uint, title string, level string, completed bool, dueDate string) error {
	// 先查询数据
	todo, err := a.GetTodo(id)
	if err != nil {
		slog.Error("更新待办事项失败", "错误信息", err.Error())
		return err
	}

	// 更新数据
	todo.Title = title
	todo.Level = level
	todo.Completed = completed

	if dueDate != "" {
		parsedTime, err := time.Parse("2006-01-02T15:04", dueDate)
		if err != nil {
			return fmt.Errorf("invalid date format: %v", err)
		}
		todo.DueDate = parsedTime
	}

	todo.Id = id
	return db.UpdateTodo(todo)
}

// GetTodo 获取指定ID的待办事项
func (a *App) GetTodo(id uint) (*db.Todo, error) {
	var todo db.Todo
	err := db.GetDB().First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// DeleteTodo 删除待办事项
func (a *App) DeleteTodo(id uint) error {
	return db.DeleteTodo(id)
}

// GetAllTodos 获取所有待办事项
func (a *App) GetAllTodos(filter string, date string, state int) ([]db.Todo, error) {
	return db.GetTodos(filter, date, state)
}
