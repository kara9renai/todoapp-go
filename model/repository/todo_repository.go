package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kara9renai/todoapp-go/model/entity"
)

type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
}

type todoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	const query = `SELECT id, title, content FROM todo ORDER BY id DESC`
	todos = []entity.TodoEntity{}
	// クエリの実行
	rows, err := Db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		todo := entity.TodoEntity{}
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Content)
		if err != nil {
			log.Print(err)
			return
		}
		todos = append(todos, todo)
	}
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
