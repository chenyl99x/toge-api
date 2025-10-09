package service

import (
	"context"
	"git.lulumia.fun/root/toge-api/internal/domain"
	"git.lulumia.fun/root/toge-api/internal/model"
	"git.lulumia.fun/root/toge-api/pkg/logger"
	"git.lulumia.fun/root/toge-api/pkg/pagination"
)

type personService struct {
	repo domain.PersonRepository
}

func NewPersonService(repo domain.PersonRepository) domain.PersonService {
	return &personService{repo: repo}
}

func (s *personService) Create(ctx context.Context, person *model.Person) error {
	if err := s.repo.Create(ctx, person); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to create person", "error", err.Error())
		return err
	}

	logger.InfoWithTrace(ctx, "Person created successfully", "person_id", person.ID)
	return nil
}

func (s *personService) GetByID(ctx context.Context, id uint) (*model.Person, error) {
	person, err := s.repo.GetByID(ctx, id)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get person by ID", "error", err.Error(), "person_id", id)
		return nil, err
	}
	logger.InfoWithTrace(ctx, "Person retrieved by ID", "person_id", id)
	return person, nil
}

func (s *personService) GetAll(ctx context.Context) ([]model.Person, error) {
	persons, err := s.repo.GetAll(ctx)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get all persons", "error", err.Error())
		return nil, err
	}
	logger.InfoWithTrace(ctx, "All persons retrieved", "count", len(persons))
	return persons, nil
}

func (s *personService) GetAllWithPagination(ctx context.Context, page *pagination.PageRequest) (*pagination.PageResponse, error) {
	persons, total, err := s.repo.GetAllWithPagination(ctx, page)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get persons with pagination", "error", err.Error(), "page", page.Page, "pageSize", page.PageSize)
		return nil, err
	}

	pageResponse := pagination.NewPageResponse(persons, total, page.Page, page.PageSize)
	logger.InfoWithTrace(ctx, "Persons retrieved with pagination", "count", len(persons), "total", total, "page", page.Page, "pageSize", page.PageSize)
	return pageResponse, nil
}

func (s *personService) Update(ctx context.Context, person *model.Person) error {
	// 检查是否存在
	_, err := s.repo.GetByID(ctx, person.ID)
	if err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get existing person for update", "error", err.Error(), "person_id", person.ID)
		return err
	}

	if err := s.repo.Update(ctx, person); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to update person", "error", err.Error(), "person_id", person.ID)
		return err
	}

	logger.InfoWithTrace(ctx, "Person updated successfully", "person_id", person.ID)
	return nil
}

func (s *personService) Delete(ctx context.Context, id uint) error {
	// 检查是否存在
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to get person for deletion", "error", err.Error(), "person_id", id)
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		logger.ErrorWithTrace(ctx, "Failed to delete person", "error", err.Error(), "person_id", id)
		return err
	}
	logger.InfoWithTrace(ctx, "Person deleted successfully", "person_id", id)
	return nil
}
