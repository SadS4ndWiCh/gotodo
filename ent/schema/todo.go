package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique(),
		field.String("title").
			MaxLen(64).
			NotEmpty(),
		field.Bool("completed").
			Default(false),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
