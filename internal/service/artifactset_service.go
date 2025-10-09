package service

import (
	"context"
	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/logger"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

type artifactsetService struct {
	repo domain.ArtifactSetRepository
}

func NewArtifactSetService(repo domain.ArtifactSetRepository) domain.ArtifactSetService {
	return &artifactsetService{repo: repo}
}

func (s *artifactsetService) Create(ctx context.Context, artifactset *model.ArtifactSet) error {
	if err := s.repo.Create(ctx, artifactset); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to create artifactset", "error", err.Error())
		return err
	}

	logger.InfoWithTrace(ctx, "ArtifactSet created successfully", "artifactset_id", artifactset.ID)
	return nil
}

func (s *artifactsetService) GetByID(ctx context.Context, id uint) (*model.ArtifactSet, error) {
	artifactset, err := s.repo.GetByID(ctx, id)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get artifactset by ID", "error", err.Error(), "artifactset_id", id)
		return nil, err
	}
	logger.InfoWithTrace(ctx, "ArtifactSet retrieved by ID", "artifactset_id", id)
	return artifactset, nil
}

func (s *artifactsetService) GetAll(ctx context.Context) ([]model.ArtifactSet, error) {
	artifactsets, err := s.repo.GetAll(ctx)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get all artifactsets", "error", err.Error())
		return nil, err
	}
	logger.InfoWithTrace(ctx, "All artifactsets retrieved", "count", len(artifactsets))
	return artifactsets, nil
}

func (s *artifactsetService) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error) {
	artifactsets, total, err := s.repo.GetAllWithPagination(ctx, page)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get artifactsets with pagination", "error", err.Error(), "page", page.Page, "pageSize", page.PageSize)
		return nil, err
	}

	pageResponse := pagination.NewPageResponse(artifactsets, total, page.Page, page.PageSize)
	logger.InfoWithTrace(ctx, "ArtifactSets retrieved with pagination", "count", len(artifactsets), "total", total, "page", page.Page, "pageSize", page.PageSize)
	return pageResponse, nil
}

func (s *artifactsetService) Update(ctx context.Context, artifactset *model.ArtifactSet) error {
	// 检查是否存在
	_, err := s.repo.GetByID(ctx, artifactset.ID)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get existing artifactset for update", "error", err.Error(), "artifactset_id", artifactset.ID)
		return err
	}

	if err := s.repo.Update(ctx, artifactset); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to update artifactset", "error", err.Error(), "artifactset_id", artifactset.ID)
		return err
	}

	logger.InfoWithTrace(ctx, "ArtifactSet updated successfully", "artifactset_id", artifactset.ID)
	return nil
}

func (s *artifactsetService) Delete(ctx context.Context, id uint) error {
	// 检查是否存在
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get artifactset for deletion", "error", err.Error(), "artifactset_id", id)
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to delete artifactset", "error", err.Error(), "artifactset_id", id)
		return err
	}
	logger.InfoWithTrace(ctx, "ArtifactSet deleted successfully", "artifactset_id", id)
	return nil
}
