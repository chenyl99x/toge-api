package domain

import (
	"context"

	"git.lulumia.fun/root/toge-api/internal/consts"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

// ArtifactRepository Artifact 仓库接口
type ArtifactRepository interface {
	Create(ctx context.Context, artifact *model.Artifact) error
	GetByID(ctx context.Context, id uint) (*model.Artifact, error)
	GetAll(ctx context.Context) ([]model.Artifact, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.Artifact, int64, error)
	Update(ctx context.Context, artifact *model.Artifact) error
	Delete(ctx context.Context, id uint) error
}

// ArtifactService Artifact 服务接口
type ArtifactService interface {
	Create(ctx context.Context, artifact *model.Artifact) error
	GetByID(ctx context.Context, id uint) (*model.Artifact, error)
	GetAll(ctx context.Context) ([]model.Artifact, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error)
	Update(ctx context.Context, artifact *model.Artifact) error
	Delete(ctx context.Context, id uint) error
}

// CreateArtifactRequest 创建 Artifact 请求
type CreateArtifactRequest struct {
	ArtifactSetID uint `json:"artifact_set_id"`

	Name string `json:"name"`

	Type consts.ArtifactType `json:"type"`

	Description string `json:"description"`

	Story string `json:"story"`
}

// UpdateArtifactRequest 更新 Artifact 请求
type UpdateArtifactRequest struct {
	ArtifactSetID *uint `json:"artifact_set_id"`

	Name *string `json:"name"`

	Type *consts.ArtifactType `json:"type"`

	Description *string `json:"description"`

	Story *string `json:"story"`
}
