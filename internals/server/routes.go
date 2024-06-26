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
	api.Handle("GET /todos/{id}", http.HandlerFunc(todoController.GetOneTodo))
	api.Handle("POST /todos/{id}/complete", http.HandlerFunc(todoController.ToggleTodo))

	return api
}
