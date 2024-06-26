package requests

import "github.com/google/uuid"

type CreateTodoRequest struct {
	Title string `json:"title" binding:"required,max=64"`
}

type GetOneTodoRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}

type ToggleCompleteTodoRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}

type DeleteTodoRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
