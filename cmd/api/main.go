package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kara9renai/todoapp-go/controller"
	"github.com/kara9renai/todoapp-go/model/repository"
)

var (
	tr = repository.NewTodoRepository()
	tc = controller.NewTodoController(tr)
	ro = controller.NewRouter(tc)
)

func main() {
	const defaultPort = ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/todos/", ro.HandleTodosRequest)

	fmt.Println("server start ... http://localhost" + defaultPort)

	log.Fatal(http.ListenAndServe(defaultPort, mux))
}
