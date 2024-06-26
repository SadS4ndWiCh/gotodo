# 🔖 gotodo

Learning to use [ent](https://entgo.io/) ORM with GO.

## 🍞 Bake

| Method | Route                 | Description                    |
| ------ | --------------------- | ------------------------------ |
| GET    | `/todos`              | Get all todos                  |
| GET    | `/todos/:id`          | Get a single todo              |
| POST   | `/todos`              | Create a new todo              |
| DELETE | `/todos/:id`          | Delete an existing todo        |
| POST   | `/todos/:id/complete` | Toggle todo's completed status |