package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kara9renai/todoapp-go/model/entity"
)

type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
	InsertTodo(todo entity.TodoEntity) (id int, err error)
	UpdateTodo(todo entity.TodoEntity) (err error)
	DeleteTodo(id int) (err error)
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
	const insert = `INSERT INTO todo (title, content) VALUES (?, ?)`

	// prepare statement
	stmt, err := Db.Prepare(insert)
	if err != nil {
		log.Println(err)
		return
	}

	result, err := stmt.Exec(todo.Title, todo.Content)
	if err != nil {
		log.Println(err)
		return
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}
	id = int(lastId)

	return
}

func (tr *todoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	const update = `UPDATE todo SET title = ?, content = ? WHERE id = ?`

	_, err = Db.Exec(update, todo.Title, todo.Content, todo.ID)
	return
}

func (tr *todoRepository) DeleteTodo(id int) (err error) {
	const deletefmt = `DELETE FROM todo WHERE id = ?`

	_, err = Db.Exec(deletefmt, id)
	return
}
