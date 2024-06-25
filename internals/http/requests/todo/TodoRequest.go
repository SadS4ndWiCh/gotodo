package requests

type CreateTodoRequest struct {
	Title string `json:"title" binding:"required,max=64"`
}
