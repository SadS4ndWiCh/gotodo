package usecases

import (
	"context"

	"github.com/SadS4ndWiCh/gotodo/ent/todo"
	"github.com/SadS4ndWiCh/gotodo/internals/database"
	requests "github.com/SadS4ndWiCh/gotodo/internals/http/requests/todo"
)

func (TodoUseCase) DeleteTodoAction(req requests.DeleteTodoRequest) error {
	client := database.GetClient()
	ctx := context.Background()

	t, err := client.Todo.
		Query().
		Where(
			todo.ID(req.Id),
		).
		Only(ctx)

	if err != nil {
		return err
	}

	err = client.Todo.
		DeleteOne(t).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
