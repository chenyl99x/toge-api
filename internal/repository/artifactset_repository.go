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

type artifactsetRepository struct{}

func NewArtifactSetRepository() domain.ArtifactSetRepository {
	return &artifactsetRepository{}
}

func (r *artifactsetRepository) Create(ctx context.Context, artifactset *model.ArtifactSet) error {
	return database.DB.WithContext(ctx).Create(artifactset).Error
}

func (r *artifactsetRepository) GetByID(ctx context.Context, id uint) (*model.ArtifactSet, error) {
	var artifactset model.ArtifactSet
	err := database.DB.WithContext(ctx).First(&artifactset, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("artifactset not found")
		}
		return nil, err
	}
	return &artifactset, nil
}

func (r *artifactsetRepository) GetAll(ctx context.Context) ([]model.ArtifactSet, error) {
	var artifactsets []model.ArtifactSet
	err := database.DB.WithContext(ctx).Find(&artifactsets).Error
	return artifactsets, err
}

func (r *artifactsetRepository) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.ArtifactSet, int64, error) {
	var artifactsets []model.ArtifactSet
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
	if err := query.Model(&model.ArtifactSet{}).Count(&total).Error; err != nil {
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
	err := query.Offset(offset).Limit(limit).Find(&artifactsets).Error

	return artifactsets, total, err
}

func (r *artifactsetRepository) Update(ctx context.Context, artifactset *model.ArtifactSet) error {
	// 检查记录是否存在
	var existingArtifactSet model.ArtifactSet
	if err := database.DB.WithContext(ctx).First(&existingArtifactSet, artifactset.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("artifactset not found")
		}
		return err
	}

	return database.DB.WithContext(ctx).Save(artifactset).Error
}

func (r *artifactsetRepository) Delete(ctx context.Context, id uint) error {
	// 检查记录是否存在
	var artifactset model.ArtifactSet
	if err := database.DB.WithContext(ctx).First(&artifactset, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("artifactset not found")
		}
		return err
	}

	return database.DB.WithContext(ctx).Delete(&artifactset).Error
}
