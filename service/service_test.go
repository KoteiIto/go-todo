package service

import (
	"testing"

	"github.com/KoteiIto/go-todo/entity"
	"github.com/stretchr/testify/assert"

	"github.com/KoteiIto/go-todo/repository"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
)

func TestGetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := repository.NewMockRepository(ctrl)
	mockRepository.EXPECT().GetList(2, 0).Return(entity.TodoList{
		entity.Todo{
			ID:          1,
			Title:       "title",
			Description: "description",
		},
		entity.Todo{
			ID:          2,
			Title:       "title",
			Description: "description",
		},
	}, nil)
	service := NewServiceImpl(2, mockRepository)

	cases := []struct {
		input  int
		expect entity.TodoList
		err    error
	}{
		{
			input: 1,
			expect: entity.TodoList{
				entity.Todo{
					ID:          1,
					Title:       "title",
					Description: "description",
				},
				entity.Todo{
					ID:          2,
					Title:       "title",
					Description: "description",
				},
			},
			err: nil,
		},
		{
			input:  0,
			expect: nil,
			err:    errors.New("invalid page 0"),
		},
	}
	for _, c := range cases {
		val, err := service.GetList(c.input)
		if c.err != nil {
			assert.Error(t, err)
			assert.Equal(t, c.err.Error(), err.Error())
		}
		assert.Equal(t, c.expect, val)
	}
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := repository.NewMockRepository(ctrl)
	mockRepository.EXPECT().Get(1).Return(1, &entity.Todo{
		ID:          1,
		Title:       "title",
		Description: "description",
	}, nil)
	mockRepository.EXPECT().Get(1000).Return(-1, nil, errors.New("invalid id 1000"))
	service := NewServiceImpl(10, mockRepository)

	cases := []struct {
		input  int
		expect *entity.Todo
		err    error
	}{
		{
			input: 1,
			expect: &entity.Todo{
				ID:          1,
				Title:       "title",
				Description: "description",
			},
			err: nil,
		},
		{
			input:  1000,
			expect: nil,
			err:    errors.New("invalid id 1000"),
		},
	}
	for _, c := range cases {
		val, err := service.Get(c.input)
		if c.err != nil {
			assert.Error(t, err)
			assert.Equal(t, c.err.Error(), err.Error())
		}
		assert.Equal(t, c.expect, val)
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	todo := entity.Todo{}
	mockRepository := repository.NewMockRepository(ctrl)
	mockRepository.EXPECT().Create(&todo).Return(nil)
	service := NewServiceImpl(10, mockRepository)

	err := service.Create(&todo)
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	todo := entity.Todo{}
	mockRepository := repository.NewMockRepository(ctrl)
	mockRepository.EXPECT().Update(&todo).Return(nil)
	service := NewServiceImpl(10, mockRepository)

	err := service.Update(&todo)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := repository.NewMockRepository(ctrl)
	mockRepository.EXPECT().Delete(1).Return(nil)
	service := NewServiceImpl(10, mockRepository)

	err := service.Delete(1)
	assert.Nil(t, err)
}
