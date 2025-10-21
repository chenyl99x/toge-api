package service

import (
	"context"

	"github.com/chenyl99x/toge-api/internal/domain"
	"github.com/chenyl99x/toge-api/internal/model"
	"github.com/chenyl99x/toge-api/pkg/logger"
	"github.com/chenyl99x/toge-api/pkg/pagination"
)

type spaceService struct {
	repo domain.SpaceRepository
}

func (s spaceService) Create(ctx context.Context, space *model.Space) error {
	if err := s.repo.Create(ctx, space); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to create space", "error", err.Error(), "name", space.Name)
		return err
	}

	logger.InfoWithTrace(ctx, "space created successfully", "id", space.ID, "name", space.Name)
	return nil
}

func (s spaceService) GetByID(ctx context.Context, id uint) (*model.Space, error) {
	space, err := s.repo.GetByID(ctx, id)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get space by ID", "error", err.Error(), "id", id)
		return nil, err
	}
	logger.InfoWithTrace(ctx, "space retrieved by ID", "id", id)
	return space, nil
}

func (s spaceService) GetAll(ctx context.Context) ([]model.Space, error) {
	spaces, err := s.repo.GetAll(ctx)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get all spaces", "error", err.Error())
		return nil, err
	}
	logger.InfoWithTrace(ctx, "All spaces retrieved", "count", len(spaces))
	return spaces, nil
}

func (s spaceService) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error) {
	spaces, total, err := s.repo.GetAllWithPagination(ctx, page)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get spaces with pagination", "error", err.Error(), "page", page.Page, "pageSize", page.PageSize)
		return nil, err
	}

	pageResponse := pagination.NewPageResponse(spaces, total, page.Page, page.PageSize)
	logger.InfoWithTrace(ctx, "spaces retrieved with pagination", "count", len(spaces), "total", total, "page", page.Page, "pageSize", page.PageSize)
	return pageResponse, nil
}

func (s spaceService) Update(ctx context.Context, space *model.Space) error {
	if err := s.repo.Update(ctx, space); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to update space", "error", err.Error(), "id", space.ID)
		return err
	}

	logger.InfoWithTrace(ctx, "space updated successfully", "space_id", space.ID, "name", space.Name)
	return nil
}

func (s spaceService) Delete(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to delete space", "error", err.Error(), "id", id)
		return err
	}
	logger.InfoWithTrace(ctx, "space deleted successfully", "id", id)
	return nil
}

func NewSpaceService(repo domain.SpaceRepository) domain.SpaceService {
	return &spaceService{repo: repo}
}
