package db

// 统计信息
type Stat struct {
	AllCount         int `json:"all_count"`         // 所有待办事项数量
	UncompletedCount int `json:"uncompleted_count"` // 未完成待办事项数量
	CompletedCount   int `json:"completed_count"`   // 已完成待办事项数量
}

type TodoList struct {
	Stat
	List []Todo `json:"list"` // 待办事项列表
}

type QryTodoList struct {
	Filter string `form:"filter" json:"filter"` // 过滤条件
	State  int    `form:"state" json:"state"`   // 状态
	Date   string `form:"date" json:"date"`     // 日期
}
