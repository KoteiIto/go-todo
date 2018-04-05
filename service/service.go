package service

import (
	"github.com/KoteiIto/go-todo/entity"
	"github.com/KoteiIto/go-todo/repository"
	"github.com/pkg/errors"
)

type Service interface {
	GetList(page int) (entity.TodoList, error)
	Get(id int) (*entity.Todo, error)
	Create(new *entity.Todo) error
	Update(update *entity.Todo) error
	Delete(id int) error
}

type ServiceImpl struct {
	limit      int
	repository repository.Repository
}

func NewServiceImpl(limit int, repository repository.Repository) *ServiceImpl {
	return &ServiceImpl{
		limit:      limit,
		repository: repository,
	}
}

func (s *ServiceImpl) GetList(page int) (entity.TodoList, error) {
	if page <= 0 {
		return nil, errors.Errorf("invalid page %d", page)
	}

	offset := (page - 1) * s.limit
	return s.repository.GetList(s.limit, offset)
}

func (s *ServiceImpl) Get(id int) (*entity.Todo, error) {
	_, todo, err := s.repository.Get(id)
	return todo, err
}

func (s *ServiceImpl) Create(new *entity.Todo) error {
	return s.repository.Create(new)
}

func (s *ServiceImpl) Update(update *entity.Todo) error {
	return s.repository.Update(update)
}

func (s *ServiceImpl) Delete(id int) error {
	return s.repository.Delete(id)
}
