package part4_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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

func TestPart4(t *testing.T) {
	t.Run("TestPutTodo_OK", func(t *testing.T) {
		json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/todos/2", json)

		target := controller.NewTodoController(&mock.MockTodoRepository{})
		target.PutTodo(w, r)

		if w.Code != 204 {
			t.Errorf("Response code is %v", w.Code)
		}
	})

	t.Run("TestPutTodo_InvalidPath", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/todos/", nil)

		target := controller.NewTodoController(&mock.MockTodoRepository{})
		target.PutTodo(w, r)

		if w.Code != 400 {
			t.Errorf("Response code is %v", w.Code)
		}
	})

	t.Run("TestPutTodo_Error", func(t *testing.T) {
		json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/todos/2", json)

		target := controller.NewTodoController(&mock.MockTodoRepositoryError{})
		target.PutTodo(w, r)

		if w.Code != 500 {
			t.Errorf("Response code is %v", w.Code)
		}
	})

	t.Run("shouldTestPutTodoSuccessed", func(t *testing.T) {
		json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
		r, _ := http.NewRequest("PUT", "/todos/2", json)
		w := httptest.NewRecorder()

		wantCode := 204
		mux.ServeHTTP(w, r)

		if w.Code != wantCode {
			t.Errorf("Response code is got: %v want: %v", w.Code, wantCode)
		}
	})
}
