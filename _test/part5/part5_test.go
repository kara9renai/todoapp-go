package part5_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	mock "github.com/kara9renai/todoapp-go/_mock"
	"github.com/kara9renai/todoapp-go/controller"
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

func TestPart5(t *testing.T) {
	t.Run("TestDeleteTodo_OK", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("DELETE", "/todos/2", nil)

		target := controller.NewTodoController(&mock.MockTodoRepository{})
		target.DeleteTodo(w, r)

		if w.Code != 204 {
			t.Errorf("Response code is %v", w.Code)
		}
	})

	t.Run("TestDeleteTodo_InvalidPath", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("DELETE", "/todos/", nil)

		target := controller.NewTodoController(&mock.MockTodoRepositoryError{})
		target.DeleteTodo(w, r)

		if w.Code != 400 {
			t.Errorf("Response code is %v", w.Code)
		}
	})

	t.Run("TestDeleteTodo_Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("DELETE", "/todos/2", nil)

		target := controller.NewTodoController(&mock.MockTodoRepositoryError{})
		target.DeleteTodo(w, r)

		if w.Code != 500 {
			t.Errorf("Response code is %v", w.Code)
		}
	})

}
