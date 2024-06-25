package usecases

import (
	"context"
	"log"

	"github.com/SadS4ndWiCh/gotodo/ent"
	"github.com/SadS4ndWiCh/gotodo/internals/database"
	requests "github.com/SadS4ndWiCh/gotodo/internals/http/requests/todo"
	"github.com/google/uuid"
)

func (TodoUseCase) CreateTodoAction(req requests.CreateTodoRequest) (*ent.Todo, error) {
	client := database.GetClient()
	ctx := context.Background()

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	t, err := client.Todo.
		Create().
		SetID(id).
		SetTitle(req.Title).
		Save(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return t, nil
}
