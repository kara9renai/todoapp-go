package controller

import "net/http"

type Router interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	tc TodoController
}

func NewRouter(tc TodoController) Router {
	return &router{tc}
}

func (ro *router) HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		w.WriteHeader(405)
	}
}
