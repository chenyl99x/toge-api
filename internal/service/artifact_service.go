package service

import (
	"context"
	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/logger"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

type artifactService struct {
	repo domain.ArtifactRepository
}

func NewArtifactService(repo domain.ArtifactRepository) domain.ArtifactService {
	return &artifactService{repo: repo}
}

func (s *artifactService) Create(ctx context.Context, artifact *model.Artifact) error {
	if err := s.repo.Create(ctx, artifact); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to create artifact", "error", err.Error())
		return err
	}

	logger.InfoWithTrace(ctx, "Artifact created successfully", "artifact_id", artifact.ID)
	return nil
}

func (s *artifactService) GetByID(ctx context.Context, id uint) (*model.Artifact, error) {
	artifact, err := s.repo.GetByID(ctx, id)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get artifact by ID", "error", err.Error(), "artifact_id", id)
		return nil, err
	}
	logger.InfoWithTrace(ctx, "Artifact retrieved by ID", "artifact_id", id)
	return artifact, nil
}

func (s *artifactService) GetAll(ctx context.Context) ([]model.Artifact, error) {
	artifacts, err := s.repo.GetAll(ctx)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get all artifacts", "error", err.Error())
		return nil, err
	}
	logger.InfoWithTrace(ctx, "All artifacts retrieved", "count", len(artifacts))
	return artifacts, nil
}

func (s *artifactService) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error) {
	artifacts, total, err := s.repo.GetAllWithPagination(ctx, page)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get artifacts with pagination", "error", err.Error(), "page", page.Page, "pageSize", page.PageSize)
		return nil, err
	}

	pageResponse := pagination.NewPageResponse(artifacts, total, page.Page, page.PageSize)
	logger.InfoWithTrace(ctx, "Artifacts retrieved with pagination", "count", len(artifacts), "total", total, "page", page.Page, "pageSize", page.PageSize)
	return pageResponse, nil
}

func (s *artifactService) Update(ctx context.Context, artifact *model.Artifact) error {
	// 检查是否存在
	_, err := s.repo.GetByID(ctx, artifact.ID)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get existing artifact for update", "error", err.Error(), "artifact_id", artifact.ID)
		return err
	}

	if err := s.repo.Update(ctx, artifact); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to update artifact", "error", err.Error(), "artifact_id", artifact.ID)
		return err
	}

	logger.InfoWithTrace(ctx, "Artifact updated successfully", "artifact_id", artifact.ID)
	return nil
}

func (s *artifactService) Delete(ctx context.Context, id uint) error {
	// 检查是否存在
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get artifact for deletion", "error", err.Error(), "artifact_id", id)
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to delete artifact", "error", err.Error(), "artifact_id", id)
		return err
	}
	logger.InfoWithTrace(ctx, "Artifact deleted successfully", "artifact_id", id)
	return nil
}
