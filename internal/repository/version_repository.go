package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/database"
	"git.lulumia.fun/root/toge-api/pkg/pagination"

	"gorm.io/gorm"
)

type versionRepository struct{}

func NewVersionRepository() domain.VersionRepository {
	return &versionRepository{}
}

func (r *versionRepository) Create(ctx context.Context, version *model.Version) error {
	return database.DB.WithContext(ctx).Create(version).Error
}

func (r *versionRepository) GetByID(ctx context.Context, id uint) (*model.Version, error) {
	var version model.Version
	err := database.DB.WithContext(ctx).First(&version, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("version not found")
		}
		return nil, err
	}
	return &version, nil
}

func (r *versionRepository) GetAll(ctx context.Context) ([]model.Version, error) {
	var versions []model.Version
	err := database.DB.WithContext(ctx).Find(&versions).Error
	return versions, err
}

func (r *versionRepository) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.Version, int64, error) {
	var versions []model.Version
	var total int64

	// 构建查询
	query := database.DB.WithContext(ctx)

	// 添加搜索条件
	if page.HasSearch() {
		// 验证搜索字段
		allowedSearchFields := []string{"Name"}
		if !page.ValidateSearchField(allowedSearchFields) {
			return nil, 0, fmt.Errorf("invalid search field: %s", page.GetSearchBy())
		}

		// 如果指定了搜索字段，使用该字段进行搜索
		if page.GetSearchBy() != "" {
			searchField := page.GetSearchBy()
			keyword := page.GetKeyword()
			query = query.Where(fmt.Sprintf("%s LIKE ?", searchField), "%"+keyword+"%")
		} else {
			// 如果没有指定搜索字段，在所有可搜索字段中搜索
			keyword := page.GetKeyword()
			var searchConditions []string
			var searchArgs []interface{}

			searchConditions = append(searchConditions, "Name LIKE ?")
			searchArgs = append(searchArgs, "%"+keyword+"%")

			query = query.Where(strings.Join(searchConditions, " OR "), searchArgs...)
		}
	}

	// 获取总记录数
	if err := query.Model(&model.Version{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 添加排序
	if page.HasSort() {
		// 验证排序字段
		allowedFields := []string{"ID", "Name"}
		if !page.ValidateSortField(allowedFields) {
			return nil, 0, fmt.Errorf("invalid sort field: %s", page.GetSortBy())
		}

		// 构建排序语句
		sortClause := page.GetSortBy()
		if page.GetSortOrder() == "desc" {
			sortClause += " DESC"
		} else {
			sortClause += " ASC"
		}
		query = query.Order(sortClause)
	} else {
		// 默认按创建时间倒序
		query = query.Order("created_at DESC")
	}

	// 获取分页数据
	offset := page.GetOffset()
	limit := page.GetLimit()
	err := query.Offset(offset).Limit(limit).Find(&versions).Error

	return versions, total, err
}

func (r *versionRepository) Update(ctx context.Context, version *model.Version) error {
	// 检查记录是否存在
	var existingVersion model.Version
	if err := database.DB.WithContext(ctx).First(&existingVersion, version.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("version not found")
		}
		return err
	}

	return database.DB.WithContext(ctx).Save(version).Error
}

func (r *versionRepository) Delete(ctx context.Context, id uint) error {
	// 检查记录是否存在
	var version model.Version
	if err := database.DB.WithContext(ctx).First(&version, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("version not found")
		}
		return err
	}

	return database.DB.WithContext(ctx).Delete(&version).Error
}
