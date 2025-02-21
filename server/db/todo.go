package db

import (
	"fmt"
	"time"
)

// Todo 待办事项模型
type Todo struct {
	Id             uint      `gorm:"column:id;primaryKey;autoIncrement;comment:编码;" json:"id"`                  // 编码
	Title          string    `gorm:"column:title;not null;comment:标题;" json:"title"`                            // 标题
	Level          string    `gorm:"column:level;comment:等级;" json:"level"`                                     // 等级
	Description    string    `gorm:"column:description;default:'';comment:描述;" json:"description"`              // 描述
	Completed      bool      `gorm:"column:completed;default:false;comment:完成状态;" json:"completed"`             // 完成状态
	EstimatedHours float64   `gorm:"column:estimated_hours;default:0;comment:预计耗时(小时);" json:"estimated_hours"` // 预计耗时（小时）
	Progress       int       `gorm:"column:progress;default:0;comment:完成进度;" json:"progress"`                   // 完成进度（0-100）
	DueDate        time.Time `gorm:"column:due_date;comment:开始时间;" json:"due_date"`                             // 添加截止日期字段
	GroupId        uint      `gorm:"column:group_id;comment:分组ID;" json:"group_id"`                             // 分组ID
	Group          TodoGroup `gorm:"foreignKey:GroupId;references:Id;comment:所属分组;" json:"group"`               // 所属分组
	CreatedAt      time.Time `gorm:"column:created_at;comment:创建时间;" json:"created_at"`                         // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at;comment:更新时间;" json:"updated_at"`                         // 更新时间
}

type TodoService struct{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

// CreateTodo 创建待办事项
func (srv *TodoService) CreateTodo(todo *Todo) error {
	return GetDB().Create(todo).Error
}

// GetTodos 获取所有待办事项
func (srv *TodoService) GetTodoList(condition QryTodoList) ([]Todo, error) {
	var todos []Todo

	qry := GetDB().Order("due_date asc")
	qry = qry.Where("substring(due_date, 1, 10) = ?", condition.Date)

	qry = qry.Where("group_id =?", condition.GroupId)

	if condition.Filter != "" {
		qry = qry.Where("title like ?", fmt.Sprintf(`%%%s%%`, condition.Filter))
	}

	if condition.State != -1 {
		qry = qry.Where("completed = ?", condition.State)
	}

	err := qry.Find(&todos).Error
	return todos, err
}

func (srv *TodoService) GetDateTodoStat(groupId uint, date string) (*Stat, error) {
	qry := GetDB().Order("due_date asc")
	qry = qry.Where("substring(due_date, 1, 10) = ?", date)
	cols := `sum(case when completed = 1 then 1 else 0 end) as completed_count, 
	sum(case when completed = 0 then 1 else 0 end) as uncompleted_count, 
	count(*) as all_count,
	group_id`
	stat := Stat{}
	err := qry.Select(cols).Model(&Todo{}).Group("group_id").Where("group_id = ?", groupId).Find(&stat).Error
	return &stat, err
}

// GetTodosByDate 获取指定日期的待办事项
func (srv *TodoService) GetTodosByDate(date time.Time) ([]Todo, error) {
	var todos []Todo
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := GetDB().Where("due_date BETWEEN ? AND ?", startOfDay, endOfDay).Order("due_date asc").Find(&todos).Error
	return todos, err
}

// UpdateTodo 更新待办事项
func (srv *TodoService) UpdateTodo(todo *Todo) error {
	return GetDB().Save(todo).Error
}

// DeleteTodo 删除待办事项
func (srv *TodoService) DeleteTodo(id uint) error {
	return GetDB().Delete(&Todo{}, id).Error
}
