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

type nationRepository struct{}

func NewNationRepository() domain.NationRepository {
	return &nationRepository{}
}

func (r *nationRepository) Create(ctx context.Context, nation *model.Nation) error {
	return database.DB.WithContext(ctx).Create(nation).Error
}

func (r *nationRepository) GetByID(ctx context.Context, id uint) (*model.Nation, error) {
	var nation model.Nation
	err := database.DB.WithContext(ctx).First(&nation, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("nation not found")
		}
		return nil, err
	}
	return &nation, nil
}

func (r *nationRepository) GetAll(ctx context.Context) ([]model.Nation, error) {
	var nations []model.Nation
	err := database.DB.WithContext(ctx).Find(&nations).Error
	return nations, err
}

func (r *nationRepository) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.Nation, int64, error) {
	var nations []model.Nation
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
	if err := query.Model(&model.Nation{}).Count(&total).Error; err != nil {
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
	err := query.Offset(offset).Limit(limit).Find(&nations).Error

	return nations, total, err
}

func (r *nationRepository) Update(ctx context.Context, nation *model.Nation) error {
	// 检查记录是否存在
	var existingNation model.Nation
	if err := database.DB.WithContext(ctx).First(&existingNation, nation.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("nation not found")
		}
		return err
	}

	return database.DB.WithContext(ctx).Save(nation).Error
}

func (r *nationRepository) Delete(ctx context.Context, id uint) error {
	// 检查记录是否存在
	var nation model.Nation
	if err := database.DB.WithContext(ctx).First(&nation, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("nation not found")
		}
		return err
	}

	return database.DB.WithContext(ctx).Delete(&nation).Error
}
