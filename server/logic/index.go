package logic

import (
	"log/slog"
	"todo_list/server/db"
)

type Logic struct {
	TodoSrv  *db.TodoService
	GroupSrv *db.TodoGroupService
}

func NewLogic() *Logic {
	return &Logic{
		TodoSrv:  db.NewTodoService(),
		GroupSrv: db.NewTodoGroupService(),
	}
}

func (l *Logic) GetTodo(id uint) (*db.Todo, error) {
	var todo db.Todo
	err := db.GetDB().First(&todo, id).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// UpdateTodoState
//
//	@Description: 更新待办事项状态
//	@receiver l
//	@param id
//	@return error
func (l *Logic) UpdateTodoState(id uint) error {
	// 先查询数据
	todo, err := l.GetTodo(id)
	if err != nil {
		return err
	}

	// 更新数据
	todo.Completed = !todo.Completed
	return l.TodoSrv.UpdateTodo(todo)
}

// UpdateTodo
//
//	@Description: 更新待办事项
//	@receiver l
//	@param data
//	@return error
func (l *Logic) UpdateTodo(data db.Todo) error {
	// 先查询数据
	todo, err := l.GetTodo(data.Id)
	if err != nil {
		slog.Error("更新待办事项失败", "错误信息", err.Error())
		return err
	}

	// 更新数据
	todo.Title = data.Title
	todo.Level = data.Level
	todo.Completed = data.Completed
	todo.EstimatedHours = data.EstimatedHours
	todo.Progress = data.Progress

	return l.TodoSrv.UpdateTodo(todo)
}

func (l *Logic) CreateTodo(data db.Todo) error {
	slog.Info("创建待办事项", "title", data.Title, "dueDate", data.DueDate)
	return l.TodoSrv.CreateTodo(&data)
}

func (l *Logic) DeleteTodo(id uint) error {
	return l.TodoSrv.DeleteTodo(id)
}

func (l *Logic) GetTodoList(condition db.QryTodoList) (*db.TodoList, error) {
	// 获取统计信息
	stat, err := l.TodoSrv.GetDateTodoStat(condition.GroupId, condition.Date)
	if err != nil {
		slog.Error("获取待办事项统计信息失败", "错误信息", err.Error())
		return nil, err
	}

	// 获取待办事项列表
	todoList, err := l.TodoSrv.GetTodoList(condition)
	if err != nil {
		slog.Error("获取待办事项列表失败", "错误信息", err.Error())
		return nil, err
	}

	return &db.TodoList{
		Stat: *stat,
		List: todoList,
	}, nil
}

func (l *Logic) CreateTodoGroup(data db.TodoGroup) error {
	slog.Info("创建待办事项分组", "name", data.Name)
	return l.GroupSrv.CreateTodoGroup(&data)
}

func (l *Logic) UpdateTodoGroup(data db.TodoGroup) error {
	slog.Info("更新待办事项分组", "id", data.Id, "name", data.Name)
	return l.GroupSrv.UpdateTodoGroup(&data)
}

func (l *Logic) GetTodoGroup(id uint) (*db.TodoGroup, error) {
	return l.GroupSrv.GetTodoGroup(id)
}

func (l *Logic) DeleteTodoGroup(id uint) error {
	return l.GroupSrv.DeleteTodoGroup(id)
}

func (l *Logic) GetTodoGroups(filter string) ([]db.TodoGroup, error) {
	return l.GroupSrv.GetTodoGroups(filter)
}
