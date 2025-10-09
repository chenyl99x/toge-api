package model

import (
	"time"
)

// Migration 迁移记录模型
type Migration struct {
	ID          uint      `gorm:"primaryKey"`
	Version     string    `gorm:"uniqueIndex;not null;size:50"` // 迁移版本号
	Description string    `gorm:"size:255"`                     // 迁移描述
	AppliedAt   time.Time `gorm:"not null"`                     // 应用时间
	Checksum    string    `gorm:"size:64"`                      // 迁移文件校验和
}

// TableName 指定表名
func (Migration) TableName() string {
	return "migrations"
}
