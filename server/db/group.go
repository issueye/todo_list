package db

import (
	"fmt"
	"time"
)

// TodoGroup 待办事项分组模型
type TodoGroup struct {
	Id          uint      `gorm:"column:id;primaryKey;autoIncrement;comment:编码;" json:"id"`       // 编码
	Name        string    `gorm:"column:name;not null;comment:分组名称;" json:"name"`                 // 分组名称
	Description string    `gorm:"column:description;default:'';comment:分组描述;" json:"description"` // 分组描述
	CreatedAt   time.Time `gorm:"column:created_at;comment:创建时间;" json:"created_at"`              // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;comment:更新时间;" json:"updated_at"`              // 更新时间
}

type TodoGroupService struct{}

func NewTodoGroupService() *TodoGroupService {
	return &TodoGroupService{}
}

// CreateTodoGroup 创建待办事项分组
func (srv *TodoGroupService) CreateTodoGroup(group *TodoGroup) error {
	return GetDB().Create(group).Error
}

// GetTodoGroups 获取所有待办事项分组
func (srv *TodoGroupService) GetTodoGroups(filter string) ([]TodoGroup, error) {
	var groups []TodoGroup
	qry := GetDB().Order("created_at desc")

	if filter != "" {
		qry = qry.Where("name like ?", fmt.Sprintf(`%%%s%%`, filter))
	}

	err := qry.Find(&groups).Error
	return groups, err
}

// GetTodoGroup 获取指定ID的待办事项分组
func (srv *TodoGroupService) GetTodoGroup(id uint) (*TodoGroup, error) {
	var group TodoGroup
	err := GetDB().First(&group, id).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// UpdateTodoGroup 更新待办事项分组
func (srv *TodoGroupService) UpdateTodoGroup(group *TodoGroup) error {
	return GetDB().Save(group).Error
}

// DeleteTodoGroup 删除待办事项分组
func (srv *TodoGroupService) DeleteTodoGroup(id uint) error {
	return GetDB().Delete(&TodoGroup{}, id).Error
}
