package domain

import (
	"context"

	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

// NationRepository Nation 仓库接口
type NationRepository interface {
	Create(ctx context.Context, nation *model.Nation) error
	GetByID(ctx context.Context, id uint) (*model.Nation, error)
	GetAll(ctx context.Context) ([]model.Nation, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.Nation, int64, error)
	Update(ctx context.Context, nation *model.Nation) error
	Delete(ctx context.Context, id uint) error
}

// NationService Nation 服务接口
type NationService interface {
	Create(ctx context.Context, nation *model.Nation) error
	GetByID(ctx context.Context, id uint) (*model.Nation, error)
	GetAll(ctx context.Context) ([]model.Nation, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error)
	Update(ctx context.Context, nation *model.Nation) error
	Delete(ctx context.Context, id uint) error
}

// CreateNationRequest 创建 Nation 请求
type CreateNationRequest struct {
	Name string `json:"name"`
}

// UpdateNationRequest 更新 Nation 请求
type UpdateNationRequest struct {
	Name *string `json:"name"`
}
