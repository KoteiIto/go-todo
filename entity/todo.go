package entity

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Expire      time.Time `json:"expire"`
}

type TodoList []Todo
