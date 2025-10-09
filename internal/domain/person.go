package domain

import (
	"context"

	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

// PersonRepository Person 仓库接口
type PersonRepository interface {
	Create(ctx context.Context, person *model.Person) error
	GetByID(ctx context.Context, id uint) (*model.Person, error)
	GetAll(ctx context.Context) ([]model.Person, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.Person, int64, error)
	Update(ctx context.Context, person *model.Person) error
	Delete(ctx context.Context, id uint) error
}

// PersonService Person 服务接口
type PersonService interface {
	Create(ctx context.Context, person *model.Person) error
	GetByID(ctx context.Context, id uint) (*model.Person, error)
	GetAll(ctx context.Context) ([]model.Person, error)
	GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error)
	Update(ctx context.Context, person *model.Person) error
	Delete(ctx context.Context, id uint) error
}

// CreatePersonRequest 创建 Person 请求
type CreatePersonRequest struct {
	Name string `json:"name"`

	Age int `json:"age"`

	Gender string `json:"gender"`

	Email string `json:"email"`

	Phone string `json:"phone"`

	Address string `json:"address"`

	Company string `json:"company"`

	Position string `json:"position"`

	Status int `json:"status"`
}

// UpdatePersonRequest 更新 Person 请求
type UpdatePersonRequest struct {
	Name *string `json:"name"`

	Age *int `json:"age"`

	Gender *string `json:"gender"`

	Email *string `json:"email"`

	Phone *string `json:"phone"`

	Address *string `json:"address"`

	Company *string `json:"company"`

	Position *string `json:"position"`

	Status *int `json:"status"`
}
