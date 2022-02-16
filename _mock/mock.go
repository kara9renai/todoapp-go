package mock

import (
	"errors"
	"net/http"

	"github.com/kara9renai/todoapp-go/model/entity"
)

type MockTodoController struct{}

func (mtc *MockTodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (mtc *MockTodoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
}

func (mtc *MockTodoController) PutTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func (mtc *MockTodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

type MockTodoRepository struct{}

func (mtr *MockTodoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	todos = []entity.TodoEntity{}
	return
}

func (mtr *MockTodoRepository) InsertTodo(todo entity.TodoEntity) (id int, err error) {
	id = 2
	return
}

func (mtr *MockTodoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	return
}

func (mtr *MockTodoRepository) DeleteTodo(id int) (err error) {
	return
}

type MockTodoRepositoryGetTodosExist struct{}

func (mtrgex *MockTodoRepositoryGetTodosExist) GetTodos() (todos []entity.TodoEntity, err error) {
	todos = []entity.TodoEntity{}
	todos = append(todos, entity.TodoEntity{ID: 1, Title: "test1", Content: "content1"})
	todos = append(todos, entity.TodoEntity{ID: 1, Title: "test2", Content: "content2"})
	return
}

func (mtrgex *MockTodoRepositoryGetTodosExist) InsertTodo(todo entity.TodoEntity) (id int, err error) {
	return
}

func (mtrgex *MockTodoRepositoryGetTodosExist) UpdateTodo(todo entity.TodoEntity) (err error) {
	return
}

func (mtrgex *MockTodoRepositoryGetTodosExist) DeleteTodo(id int) (err error) {
	return
}

type MockTodoRepositoryError struct{}

func (mtrerr *MockTodoRepositoryError) GetTodos() (todos []entity.TodoEntity, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrerr *MockTodoRepositoryError) InsertTodo(todo entity.TodoEntity) (id int, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrerr *MockTodoRepositoryError) UpdateTodo(todo entity.TodoEntity) (err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrerr *MockTodoRepositoryError) DeleteTodo(id int) (err error) {
	err = errors.New("unexpected error occurred")
	return
}
