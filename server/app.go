package server

import (
	"context"
	"fmt"
	"log/slog"
	"todo_list/server/db"
	"todo_list/server/logger"
	"todo_list/server/logic"
)

// App struct
type App struct {
	ctx       context.Context
	TodoLogic *logic.Logic
}

// NewApp creates a new App application struct
func NewApp() *App {
	// 初始化数据库
	err := db.InitDB()
	if err != nil {
		panic(fmt.Errorf("数据库初始化失败：%w", err))
	}

	app := &App{}
	app.TodoLogic = logic.NewLogic()
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	// 初始化日志系统
	if err := logger.InitLogger(); err != nil {
		fmt.Printf("初始化日志系统失败: %v\n", err)
		return
	}
	slog.Info("应用程序启动成功")
}

// shutdown is called when the app is closing
func (a *App) Shutdown(ctx context.Context) {
	if err := db.CloseDB(); err != nil {
		slog.Error("关闭数据库失败", "错误信息", err.Error())
	}
	slog.Info("应用程序关闭")
}

// CreateTodo 创建新的待办事项
func (a *App) CreateTodo(data db.Todo) error {
	return a.TodoLogic.CreateTodo(data)
}

// UpdateTodo 更新待办事项
func (a *App) UpdateTodo(data db.Todo) error {
	return a.TodoLogic.UpdateTodo(data)
}

func (a *App) UpdateTodoState(id uint) error {
	return a.TodoLogic.UpdateTodoState(id)
}

// GetTodo 获取指定ID的待办事项
func (a *App) GetTodo(id uint) (*db.Todo, error) {
	return a.TodoLogic.GetTodo(id)
}

// DeleteTodo 删除待办事项
func (a *App) DeleteTodo(id uint) error {
	return a.TodoLogic.DeleteTodo(id)
}

// GetAllTodos 获取所有待办事项
func (a *App) GetTodoList(condition db.QryTodoList) (*db.TodoList, error) {
	return a.TodoLogic.GetTodoList(condition)
}

// CreateTodoGroup 创建待办事项分组
func (a *App) CreateTodoGroup(data db.TodoGroup) error {
	slog.Info("创建待办事项分组", "name", data.Name)
	return a.TodoLogic.CreateTodoGroup(data)
}

// UpdateTodoGroup 更新待办事项分组
func (a *App) UpdateTodoGroup(data db.TodoGroup) error {
	slog.Info("更新待办事项分组", "id", data.Id, "name", data.Name)
	return a.TodoLogic.UpdateTodoGroup(data)
}

// GetTodoGroup 获取指定ID的待办事项分组
func (a *App) GetTodoGroup(id uint) (*db.TodoGroup, error) {
	return a.TodoLogic.GetTodoGroup(id)
}

// DeleteTodoGroup 删除待办事项分组
func (a *App) DeleteTodoGroup(id uint) error {
	return a.TodoLogic.DeleteTodoGroup(id)
}

// GetAllTodoGroups 获取所有待办事项分组
func (a *App) GetAllTodoGroups(filter string) ([]db.TodoGroup, error) {
	return a.TodoLogic.GetTodoGroups(filter)
}
