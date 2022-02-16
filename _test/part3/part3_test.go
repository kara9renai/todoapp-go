package part3_test

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

func TestPart3(t *testing.T) {
	t.Run("should PostTodo OK", func(t *testing.T) {
		json := strings.NewReader(`{"title": "test-title", "content": "test-content"}`)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/todos/", json)

		target := controller.NewTodoController(&mock.MockTodoRepository{})
		target.PostTodo(w, r)

		if w.Code != 201 {
			t.Errorf("Response code  is %v", w.Code)
		}
		if w.Header().Get("Location") != r.Host+r.URL.Path+"2" {
			t.Errorf("Location is %v", w.Header().Get("Location"))

		}
	})

	t.Run("should PostTodo Error", func(t *testing.T) {
		json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/todos/", json)

		target := controller.NewTodoController(&mock.MockTodoRepositoryError{})
		target.PostTodo(w, r)

		if w.Code != 500 {
			t.Errorf("Response cod is %v", w.Code)
		}
		if w.Header().Get("Location") != "" {
			t.Errorf("Location is %v", w.Header().Get("Location"))
		}
	})

	t.Run("should PostTodo OK in router.go", func(t *testing.T) {
		json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
		r, _ := http.NewRequest("POST", "/todos/", json)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, r)

		if w.Code != 201 {
			t.Errorf("Response code is %v", w.Code)
		}
	})
}
