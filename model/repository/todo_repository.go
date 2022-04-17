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
	//以下のクエリを使用してください
	//lint:ignore U1000 It's ok because this is a sample query.
	const query = `SELECT id, title, content FROM todo ORDER BY id DESC`
	return
}

func (tr *todoRepository) InsertTodo(todo entity.TodoEntity) (id int, err error) {
	//以下のクエリを使用してください
	//lint:ignore U1000 It's ok because this is a sample query.
	const insert = `INSERT INTO todo (title, content) VALUES (?, ?)`
	return
}

func (tr *todoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	//以下のクエリを使用してください
	//lint:ignore U1000 It's ok because this is a sample query.
	const update = `UPDATE todo SET title = ?, content = ? WHERE id = ?`
	return
}

func (tr *todoRepository) DeleteTodo(id int) (err error) {
	//以下のクエリを使用してください
	//lint:ignore U1000 It's ok because this is a sample query.
	const deleteFmt = `DELETE FROM todo WHERE id = ?`
	return
}
