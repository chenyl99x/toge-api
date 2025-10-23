package model

import "gorm.io/gorm"

// Space 岛屿模型
// @Description 空间信息

type Space struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null;comment:岛屿名称" json:"name" example:"我的岛屿"` // 岛屿名称
	OwnerUserID uint   `gorm:"not null;index;comment:岛屿拥有者ID" json:"owner_user_id" example:"1"`    // 岛屿拥有者ID
	Type        string `gorm:"type:varchar(50);not null;comment:岛屿类型" json:"type" example:"情侣空间"`  // 岛屿类型:情侣空间、家庭空间
	Description string `gorm:"type:text;comment:描述" json:"description" example:"这是一个美好的岛屿"`        // 描述
}

// TableName 指定表名
func (Space) TableName() string {
	return "space"
}
