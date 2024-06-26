package usecases

import (
	"context"

	"github.com/SadS4ndWiCh/gotodo/ent"
	"github.com/SadS4ndWiCh/gotodo/ent/todo"
	"github.com/SadS4ndWiCh/gotodo/internals/database"
	requests "github.com/SadS4ndWiCh/gotodo/internals/http/requests/todo"
)

func (TodoUseCase) ToggleCompleteTodoAction(req requests.ToggleCompleteTodoRequest) (*ent.Todo, error) {
	client := database.GetClient()
	ctx := context.Background()

	t, err := client.Todo.
		Query().
		Where(
			todo.ID(req.Id),
		).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	t, err = t.
		Update().
		SetCompleted(!t.Completed).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return t, nil
}
