package domain

import (
	"context"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

// VersionRepository Version 仓库接口
type VersionRepository interface {
	Create(ctx context.Context, version *model.Version) error
	GetByID(ctx context.Context, id uint) (*model.Version, error)
	GetAll(ctx context.Context) ([]model.Version, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.Version, int64, error)
	Update(ctx context.Context, version *model.Version) error
	Delete(ctx context.Context, id uint) error
}

// VersionService Version 服务接口
type VersionService interface {
	Create(ctx context.Context, version *model.Version) error
	GetByID(ctx context.Context, id uint) (*model.Version, error)
	GetAll(ctx context.Context) ([]model.Version, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error)
	Update(ctx context.Context, version *model.Version) error
	Delete(ctx context.Context, id uint) error
}

// CreateVersionRequest 创建 Version 请求
type CreateVersionRequest struct {
	Name string `json:"name"`
}

// UpdateVersionRequest 更新 Version 请求
type UpdateVersionRequest struct {
	Name *string `json:"name"`
}
