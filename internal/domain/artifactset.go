package domain

import (
	"context"

	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

// ArtifactSetRepository ArtifactSet 仓库接口
type ArtifactSetRepository interface {
	Create(ctx context.Context, artifactset *model.ArtifactSet) error
	GetByID(ctx context.Context, id uint) (*model.ArtifactSet, error)
	GetAll(ctx context.Context) ([]model.ArtifactSet, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.ArtifactSet, int64, error)
	Update(ctx context.Context, artifactset *model.ArtifactSet) error
	Delete(ctx context.Context, id uint) error
}

// ArtifactSetService ArtifactSet 服务接口
type ArtifactSetService interface {
	Create(ctx context.Context, artifactset *model.ArtifactSet) error
	GetByID(ctx context.Context, id uint) (*model.ArtifactSet, error)
	GetAll(ctx context.Context) ([]model.ArtifactSet, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error)
	Update(ctx context.Context, artifactset *model.ArtifactSet) error
	Delete(ctx context.Context, id uint) error
}

// CreateArtifactSetRequest 创建 ArtifactSet 请求
type CreateArtifactSetRequest struct {
	Name string `json:"name"`
}

// UpdateArtifactSetRequest 更新 ArtifactSet 请求
type UpdateArtifactSetRequest struct {
	Name *string `json:"name"`
}
