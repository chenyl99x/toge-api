package model

import (
	"time"

	"gorm.io/gorm"
)

type ArtifactSet struct {
	ID        uint           `json:"id" gorm:"primaryKey" example:"1"`
	Name      string         `json:"name" gorm:"not null;size:100" example:"1.0"`
	CreatedAt time.Time      `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2023-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index" swaggerignore:"true"`
}

func (ArtifactSet) TableName() string {
	return "artifact_set"
}
