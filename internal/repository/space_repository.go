package repository

import (
	"context"
	"fmt"

	"github.com/chenyl99x/toge-api/internal/domain"
	"github.com/chenyl99x/toge-api/internal/model"
	"github.com/chenyl99x/toge-api/pkg/database"
	"github.com/chenyl99x/toge-api/pkg/pagination"
)

type spaceRepository struct {
}

func (s spaceRepository) Create(ctx context.Context, space *model.Space) error {
	return database.DB.WithContext(ctx).Create(space).Error
}

func (s spaceRepository) GetByID(ctx context.Context, id uint) (*model.Space, error) {

	var space model.Space
	return &space, database.DB.WithContext(ctx).First(&space, id).Error
}

func (s spaceRepository) GetAll(ctx context.Context) ([]model.Space, error) {
	var spaces []model.Space
	return spaces, database.DB.WithContext(ctx).Find(&spaces).Error
}

func (s spaceRepository) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.Space, int64, error) {

	var spaces []model.Space
	var total int64
	query := database.DB.WithContext(ctx)

	if page.HasSort() {
		// 验证排序字段
		allowedFields := []string{"id", "name", "created_at", "updated_at"}
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
	err := query.Offset(offset).Limit(limit).Find(&spaces).Error

	return spaces, total, err
}

func (s spaceRepository) Update(ctx context.Context, Space *model.Space) error {
	return database.DB.WithContext(ctx).Updates(Space).Error
}

func (s spaceRepository) Delete(ctx context.Context, id uint) error {
	return database.DB.WithContext(ctx).Delete(&model.Space{}, id).Error
}

func NewSpaceRepository() domain.SpaceRepository {
	return &spaceRepository{}
}
