package domain

import (
	"context"

	"github.com/chenyl99x/toge-api/internal/model"
	"github.com/chenyl99x/toge-api/pkg/pagination"
)

type SpaceRepository interface {
	Create(ctx context.Context, space *model.Space) error
	GetByID(ctx context.Context, id uint) (*model.Space, error)
	GetAll(ctx context.Context) ([]model.Space, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.Space, int64, error)
	Update(ctx context.Context, space *model.Space) error
	Delete(ctx context.Context, id uint) error
}

type SpaceService interface {
	Create(ctx context.Context, space *model.Space) error
	GetByID(ctx context.Context, id uint) (*model.Space, error)
	GetAll(ctx context.Context) ([]model.Space, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error)
	Update(ctx context.Context, space *model.Space) error
	Delete(ctx context.Context, id uint) error
}

type CreateSpaceRequest struct {
	Name        string `json:"name" binding:"required" example:"我的空间"`
	Description string `json:"description" binding:"required" example:"这是一个美好的空间"`
	Type        string `json:"type" binding:"required" example:"情侣空间"`
	OwnerUserID uint   `json:"owner_user_id" binding:"required" example:"1"`
}

type UpdateSpaceRequest struct {
	Name        string `json:"name" binding:"required" example:"我的空间"`
	Description string `json:"description" binding:"required" example:"这是一个美好的空间"`
	Type        string `json:"type" binding:"required" example:"情侣空间"`
	OwnerUserID uint   `json:"owner_user_id" binding:"required" example:"1"`
}
