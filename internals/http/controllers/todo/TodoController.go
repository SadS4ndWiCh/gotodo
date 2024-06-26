package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	requests "github.com/SadS4ndWiCh/gotodo/internals/http/requests/todo"
	usecases "github.com/SadS4ndWiCh/gotodo/internals/usecases/todo"
	"github.com/SadS4ndWiCh/gotodo/internals/utils"
	"github.com/google/uuid"
)

type TodoController struct{}

func (TodoController) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	todoUseCase := usecases.TodoUseCase{}

	t, err := todoUseCase.GetAllTodoAction()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todos, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(todos)
}

func (TodoController) GetOneTodo(w http.ResponseWriter, r *http.Request) {
	id := uuid.MustParse(r.PathValue("id"))
	todoUseCase := usecases.TodoUseCase{}

	req := requests.GetOneTodoRequest{
		Id: id,
	}

	t, err := todoUseCase.GetOneTodoAction(req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	todo, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(todo)
}

func (TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo requests.CreateTodoRequest

	if err := utils.DecodeJSONBody(w, r, &newTodo); err != nil {
		var malformedRequest utils.MalformedRequest

		if errors.As(err, &malformedRequest) {
			http.Error(w, malformedRequest.Error(), malformedRequest.Status)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	todoUseCases := usecases.TodoUseCase{}
	t, err := todoUseCases.CreateTodoAction(newTodo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	todo, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(todo)
}

func (TodoController) ToggleTodo(w http.ResponseWriter, r *http.Request) {
	id := uuid.MustParse(r.PathValue("id"))
	todoUseCase := usecases.TodoUseCase{}

	req := requests.ToggleCompleteTodoRequest{
		Id: id,
	}

	t, err := todoUseCase.ToggleCompleteTodoAction(req)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := json.Marshal(t)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(todo)
}
