package service

import (
	"context"
	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/logger"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

type nationService struct {
	repo domain.NationRepository
}

func NewNationService(repo domain.NationRepository) domain.NationService {
	return &nationService{repo: repo}
}

func (s *nationService) Create(ctx context.Context, nation *model.Nation) error {
	if err := s.repo.Create(ctx, nation); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to create nation", "error", err.Error())
		return err
	}

	logger.InfoWithTrace(ctx, "Nation created successfully", "nation_id", nation.ID)
	return nil
}

func (s *nationService) GetByID(ctx context.Context, id uint) (*model.Nation, error) {
	nation, err := s.repo.GetByID(ctx, id)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get nation by ID", "error", err.Error(), "nation_id", id)
		return nil, err
	}
	logger.InfoWithTrace(ctx, "Nation retrieved by ID", "nation_id", id)
	return nation, nil
}

func (s *nationService) GetAll(ctx context.Context) ([]model.Nation, error) {
	nations, err := s.repo.GetAll(ctx)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get all nations", "error", err.Error())
		return nil, err
	}
	logger.InfoWithTrace(ctx, "All nations retrieved", "count", len(nations))
	return nations, nil
}

func (s *nationService) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error) {
	nations, total, err := s.repo.GetAllWithPagination(ctx, page)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get nations with pagination", "error", err.Error(), "page", page.Page, "pageSize", page.PageSize)
		return nil, err
	}

	pageResponse := pagination.NewPageResponse(nations, total, page.Page, page.PageSize)
	logger.InfoWithTrace(ctx, "Nations retrieved with pagination", "count", len(nations), "total", total, "page", page.Page, "pageSize", page.PageSize)
	return pageResponse, nil
}

func (s *nationService) Update(ctx context.Context, nation *model.Nation) error {
	// 检查是否存在
	_, err := s.repo.GetByID(ctx, nation.ID)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get existing nation for update", "error", err.Error(), "nation_id", nation.ID)
		return err
	}

	if err := s.repo.Update(ctx, nation); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to update nation", "error", err.Error(), "nation_id", nation.ID)
		return err
	}

	logger.InfoWithTrace(ctx, "Nation updated successfully", "nation_id", nation.ID)
	return nil
}

func (s *nationService) Delete(ctx context.Context, id uint) error {
	// 检查是否存在
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get nation for deletion", "error", err.Error(), "nation_id", id)
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to delete nation", "error", err.Error(), "nation_id", id)
		return err
	}
	logger.InfoWithTrace(ctx, "Nation deleted successfully", "nation_id", id)
	return nil
}
