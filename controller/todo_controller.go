package controller

import (
	"net/http"

	"github.com/kara9renai/todoapp-go/model/repository"
)

type TodoController interface {
}

type toDoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &toDoController{tr}
}

func (tc *toDoController) GetTodos(w http.ResponseWriter, r *http.Request) {
}

func (tc *toDoController) PostTodo(w http.ResponseWriter, r *http.Request) {
}

func (tc *toDoController) PutTodo(w http.ResponseWriter, r *http.Request) {
}

func (tc *toDoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
}
