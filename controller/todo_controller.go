package controller

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/kara9renai/todoapp-go/controller/dto"
	"github.com/kara9renai/todoapp-go/model/entity"
	"github.com/kara9renai/todoapp-go/model/repository"
)

type TodoController interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	PostTodo(w http.ResponseWriter, r *http.Request)
	PutTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type toDoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &toDoController{tr}
}

func (tc *toDoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var todoRes []dto.TodoResponse
	for _, v := range todos {
		todoRes = append(todoRes, dto.TodoResponse{ID: v.ID, Title: v.Title, Content: v.Content})
	}

	var todosRes dto.TodosResponse
	todosRes.Todos = todoRes

	output, _ := json.MarshalIndent(todosRes.Todos, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (tc *toDoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var todoReq dto.TodoRequest
	json.Unmarshal(body, &todoReq)

	todo := entity.TodoEntity{Title: todoReq.Title, Content: todoReq.Content}

	id, err := tc.tr.InsertTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func (tc *toDoController) PutTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoReq dto.TodoRequest
	json.Unmarshal(body, &todoReq)

	todo := entity.TodoEntity{ID: todoId, Title: todoReq.Title, Content: todoReq.Content}

	err = tc.tr.UpdateTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func (tc *toDoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = tc.tr.DeleteTodo(todoId)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
