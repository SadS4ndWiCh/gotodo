package usecases

import (
	"context"

	"github.com/SadS4ndWiCh/gotodo/ent"
	"github.com/SadS4ndWiCh/gotodo/internals/database"
)

func (TodoUseCase) GetAllTodoAction() ([]*ent.Todo, error) {
	client := database.GetClient()
	ctx := context.Background()

	t, err := client.Todo.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	return t, err
}
