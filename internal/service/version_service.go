package service

import (
	"context"
	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/logger"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

type versionService struct {
	repo domain.VersionRepository
}

func NewVersionService(repo domain.VersionRepository) domain.VersionService {
	return &versionService{repo: repo}
}

func (s *versionService) Create(ctx context.Context, version *model.Version) error {
	if err := s.repo.Create(ctx, version); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to create version", "error", err.Error())
		return err
	}

	logger.InfoWithTrace(ctx, "Version created successfully", "version_id", version.ID)
	return nil
}

func (s *versionService) GetByID(ctx context.Context, id uint) (*model.Version, error) {
	version, err := s.repo.GetByID(ctx, id)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get version by ID", "error", err.Error(), "version_id", id)
		return nil, err
	}
	logger.InfoWithTrace(ctx, "Version retrieved by ID", "version_id", id)
	return version, nil
}

func (s *versionService) GetAll(ctx context.Context) ([]model.Version, error) {
	versions, err := s.repo.GetAll(ctx)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get all versions", "error", err.Error())
		return nil, err
	}
	logger.InfoWithTrace(ctx, "All versions retrieved", "count", len(versions))
	return versions, nil
}

func (s *versionService) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error) {
	versions, total, err := s.repo.GetAllWithPagination(ctx, page)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get versions with pagination", "error", err.Error(), "page", page.Page, "pageSize", page.PageSize)
		return nil, err
	}

	pageResponse := pagination.NewPageResponse(versions, total, page.Page, page.PageSize)
	logger.InfoWithTrace(ctx, "Versions retrieved with pagination", "count", len(versions), "total", total, "page", page.Page, "pageSize", page.PageSize)
	return pageResponse, nil
}

func (s *versionService) Update(ctx context.Context, version *model.Version) error {
	// 检查是否存在
	_, err := s.repo.GetByID(ctx, version.ID)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get existing version for update", "error", err.Error(), "version_id", version.ID)
		return err
	}

	if err := s.repo.Update(ctx, version); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to update version", "error", err.Error(), "version_id", version.ID)
		return err
	}

	logger.InfoWithTrace(ctx, "Version updated successfully", "version_id", version.ID)
	return nil
}

func (s *versionService) Delete(ctx context.Context, id uint) error {
	// 检查是否存在
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get version for deletion", "error", err.Error(), "version_id", id)
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to delete version", "error", err.Error(), "version_id", id)
		return err
	}
	logger.InfoWithTrace(ctx, "Version deleted successfully", "version_id", id)
	return nil
}
