package model

import (
	"time"

	"gorm.io/gorm"
)

// Person 人员模型
// @Description 人员信息
type Person struct {
	ID        uint           `json:"id" gorm:"primaryKey" example:"1"`
	Name      string         `json:"name" gorm:"not null;size:100" example:"张三"`
	Age       int            `json:"age" gorm:"not null" example:"25"`
	Gender    string         `json:"gender" gorm:"size:10" example:"男"`
	Email     string         `json:"email" gorm:"uniqueIndex;size:100" example:"zhangsan@example.com"`
	Phone     string         `json:"phone" gorm:"size:20" example:"13800138000"`
	Address   string         `json:"address" gorm:"size:255" example:"北京市朝阳区"`
	Company   string         `json:"company" gorm:"size:100" example:"科技有限公司"`
	Position  string         `json:"position" gorm:"size:50" example:"软件工程师"`
	Status    int            `json:"status" gorm:"default:1" example:"1"` // 1: 在职, 0: 离职
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggerignore:"true"`
}

func (Person) TableName() string {
	return "person"
}
