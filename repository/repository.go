package repository

import (
	"github.com/KoteiIto/go-todo/entity"
	"github.com/KoteiIto/go-todo/util"
	"github.com/pkg/errors"
)

type Repository interface {
	GetList(limit int, offset int) (entity.TodoList, error)
	Get(id int) (int, *entity.Todo, error)
	Create(new *entity.Todo) error
	Update(update *entity.Todo) error
	Delete(id int) error
}

// MemoryRepository 簡易的にTodoをメモリーに保存する
// 本来なら更新処理はゴルーチンを使用して排他処理をしないといけない
type MemoryRepository struct {
	todoList  entity.TodoList
	currentID int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		todoList:  entity.TodoList{},
		currentID: 1,
	}
}

func (r *MemoryRepository) GetList(limit int, offset int) (entity.TodoList, error) {
	if limit <= 0 {
		return nil, errors.Errorf("invalid limit %d", limit)
	}

	if offset < 0 || len(r.todoList) < offset {
		return nil, errors.Errorf("invalid offset %d", offset)
	}

	end := util.Min(offset+limit, len(r.todoList))

	return r.todoList[offset:end], nil
}

func (r *MemoryRepository) Get(id int) (int, *entity.Todo, error) {
	for i, todo := range r.todoList {
		if todo.ID == id {
			return i, &todo, nil
		}
	}
	return -1, nil, errors.Errorf("invalid id %d", id)
}

func (r *MemoryRepository) Create(new *entity.Todo) error {
	if new == nil {
		return errors.New("todo is empty")
	}
	new.ID = r.currentID
	r.todoList = append(r.todoList, *new)
	r.currentID++
	return nil
}
func (r *MemoryRepository) Update(update *entity.Todo) error {
	if update == nil {
		return errors.New("todo is empty")
	}

	i, _, err := r.Get(update.ID)
	if err != nil {
		return err
	}

	r.todoList[i] = *update
	return nil
}
func (r *MemoryRepository) Delete(id int) error {
	i, _, err := r.Get(id)
	if err != nil {
		return err
	}

	l := r.todoList[0:i]
	l = append(l, r.todoList[i+1:]...)
	r.todoList = l
	return nil
}
