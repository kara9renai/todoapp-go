package part2_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	mock "github.com/kara9renai/todoapp-go/_mock"
	"github.com/kara9renai/todoapp-go/controller"
	"github.com/kara9renai/todoapp-go/controller/dto"
)

var mux *http.ServeMux

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	target := controller.NewRouter(&mock.MockTodoController{})
	mux = http.NewServeMux()
	mux.HandleFunc("/todos/", target.HandleTodosRequest)
}

func TestPart2(t *testing.T) {
	t.Run("TestGetTodos", func(t *testing.T) {
		r, _ := http.NewRequest("GET", "/todos/", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, r)

		if w.Code != 200 {
			t.Errorf("Response code is %v", w.Code)
		}
	})

	t.Run("TestGetTodos_NotFound", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/todos/", nil)
		target := controller.NewTodoController(&mock.MockTodoRepository{})
		target.GetTodos(w, r)

		if w.Code != 200 {
			t.Errorf("Reponse code is %v", w.Code)
		}
		if w.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
		}

		body := make([]byte, w.Body.Len())
		w.Body.Read(body)
		var todosResp dto.TodosResponse
		json.Unmarshal(body, &todosResp)
		if len(todosResp.Todos) != 0 {
			t.Errorf("Response is %v", todosResp.Todos)
		}
	})

	t.Run("TestGetTodos_ExistTodo", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/todos/", nil)

		target := controller.NewTodoController(&mock.MockTodoRepositoryGetTodosExist{})
		target.GetTodos(w, r)

		if w.Code != 200 {
			t.Errorf("Response cod is %v", w.Code)
		}
		if w.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
		}

		body := make([]byte, w.Body.Len())
		w.Body.Read(body)
		var todosResponse dto.TodosResponse
		json.Unmarshal(body, &todosResponse.Todos)
		if len(todosResponse.Todos) != 2 {
			t.Errorf("Response is %v", todosResponse.Todos)
		}
	})

	t.Run("TestGetTodos_Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/todos/", nil)

		target := controller.NewTodoController(&mock.MockTodoRepositoryError{})
		target.GetTodos(w, r)

		if w.Code != 500 {
			t.Errorf("Response cod is %v", w.Code)
		}
		if w.Header().Get("Content-Type") != "" {
			t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
		}

		if w.Body.Len() != 0 {
			t.Errorf("body is %v", w.Body.Len())
		}
	})
}
