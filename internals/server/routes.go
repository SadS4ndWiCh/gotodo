package server

import (
	"net/http"

	controllers "github.com/SadS4ndWiCh/gotodo/internals/http/controllers/todo"
)

func (s *Server) Bootstrap() *http.ServeMux {
	api := http.NewServeMux()

	todoController := controllers.TodoController{}
	api.Handle("GET /todos", http.HandlerFunc(todoController.GetAllTodo))
	api.Handle("POST /todos", http.HandlerFunc(todoController.CreateTodo))

	return api
}
