package model

import "gorm.io/gorm"

// Space 空间模型
// @Description 空间信息

type Space struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null;comment:空间名称" json:"name" example:"我的空间"` // 空间名称
	OwnerUserID uint   `gorm:"not null;index;comment:空间拥有者ID" json:"owner_user_id" example:"1"`    // 空间拥有者ID
	Type        string `gorm:"type:varchar(50);not null;comment:空间类型" json:"type" example:"情侣空间"`  // 空间类型:情侣空间、家庭空间
	Description string `gorm:"type:text;comment:描述" json:"description" example:"这是一个美好的空间"`        // 描述
}

// TableName 指定表名
func (Space) TableName() string {
	return "space"
}
