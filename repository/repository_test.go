package repository

import (
	"testing"

	"github.com/KoteiIto/go-todo/entity"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	todoList := entity.TodoList{
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
		entity.Todo{
			ID:          3,
			Title:       "title",
			Description: "description",
		},
	}

	type input struct {
		limit    int
		offset   int
		todoList entity.TodoList
	}
	cases := []struct {
		input  input
		expect entity.TodoList
		err    error
	}{
		{
			input: input{
				limit:    2,
				offset:   0,
				todoList: todoList,
			},
			expect: todoList[0:2],
			err:    nil,
		}, {
			input: input{
				limit:    3,
				offset:   0,
				todoList: todoList,
			},
			expect: todoList,
			err:    nil,
		},
		{
			input: input{
				limit:    1,
				offset:   1,
				todoList: todoList,
			},
			expect: todoList[1:2],
			err:    nil,
		},
		{
			input: input{
				limit:    1,
				offset:   3,
				todoList: todoList,
			},
			expect: todoList[3:3],
			err:    nil,
		},
		{
			input: input{
				limit:    0,
				offset:   1,
				todoList: todoList,
			},
			expect: nil,
			err:    errors.Errorf("invalid limit 0"),
		},
		{
			input: input{
				limit:    2,
				offset:   4,
				todoList: todoList,
			},
			expect: nil,
			err:    errors.Errorf("invalid offset 4"),
		},
	}
	for _, c := range cases {
		repository := MemoryRepository{
			todoList: c.input.todoList,
		}
		val, err := repository.GetList(c.input.limit, c.input.offset)
		if c.err != nil {
			assert.Error(t, err)
			assert.Equal(t, c.err.Error(), err.Error())
		}
		assert.Equal(t, c.expect, val)
	}
}
