package repository

import (
	"context"
	"errors"
	"fmt"
	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/database"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
	"strings"

	"gorm.io/gorm"
)

type personRepository struct{}

func NewPersonRepository() domain.PersonRepository {
	return &personRepository{}
}

func (r *personRepository) Create(ctx context.Context, person *model.Person) error {
	return database.DB.WithContext(ctx).Create(person).Error
}

func (r *personRepository) GetByID(ctx context.Context, id uint) (*model.Person, error) {
	var person model.Person
	err := database.DB.WithContext(ctx).First(&person, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("person not found")
		}
		return nil, err
	}
	return &person, nil
}

func (r *personRepository) GetAll(ctx context.Context) ([]model.Person, error) {
	var persons []model.Person
	err := database.DB.WithContext(ctx).Find(&persons).Error
	return persons, err
}

func (r *personRepository) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) ([]model.Person, int64, error) {
	var persons []model.Person
	var total int64

	// 构建查询
	query := database.DB.WithContext(ctx)

	// 添加搜索条件
	if page.HasSearch() {
		// 验证搜索字段
		allowedSearchFields := []string{"Name", "Email"}
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

			searchConditions = append(searchConditions, "Email LIKE ?")
			searchArgs = append(searchArgs, "%"+keyword+"%")

			query = query.Where(strings.Join(searchConditions, " OR "), searchArgs...)
		}
	}

	// 获取总记录数
	if err := query.Model(&model.Person{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 添加排序
	if page.HasSort() {
		// 验证排序字段
		allowedFields := []string{"ID", "Name", "Status"}
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
	err := query.Offset(offset).Limit(limit).Find(&persons).Error

	return persons, total, err
}

func (r *personRepository) Update(ctx context.Context, person *model.Person) error {
	// 检查记录是否存在
	var existingPerson model.Person
	if err := database.DB.WithContext(ctx).First(&existingPerson, person.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("person not found")
		}
		return err
	}

	return database.DB.WithContext(ctx).Save(person).Error
}

func (r *personRepository) Delete(ctx context.Context, id uint) error {
	// 检查记录是否存在
	var person model.Person
	if err := database.DB.WithContext(ctx).First(&person, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("person not found")
		}
		return err
	}

	return database.DB.WithContext(ctx).Delete(&person).Error
}
