package model

import (
	"time"

	"git.lulumia.fun/root/toge-api/internal/consts"
	"gorm.io/gorm"
)

type Artifact struct {
	ID            uint                `json:"id" gorm:"primaryKey" example:"1"`
	ArtifactSetID uint                `json:"artifact_set_id" gorm:"not null" example:"1"`
	Name          string              `json:"name" gorm:"not null;size:100" example:"1.0"`
	Type          consts.ArtifactType `json:"type" gorm:"not null" example:"生之花"`
	Description   string              `json:"description" gorm:"not null;type:text" example:"1.0"`
	Story         string              `json:"story" gorm:"not null;type:text" example:"1.0"`
	CreatedAt     time.Time           `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt     time.Time           `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt     gorm.DeletedAt      `json:"deleted_at,omitempty" gorm:"index" swaggerignore:"true"`
}

func (Artifact) TableName() string {
	return "artifact"
}
