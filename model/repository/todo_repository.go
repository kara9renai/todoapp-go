package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kara9renai/todoapp-go/model/entity"
)

type TodoRepository interface {
}

type todoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	return
}

func (tr *todoRepository) InsertTodo(todo entity.TodoEntity) (id int, err error) {
	return
}

func (tr *todoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	return
}

func (tr *todoRepository) DeleteTodo(id int) (err error) {
	return
}
