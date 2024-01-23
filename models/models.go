package models

import (
	"hellocheck/internal/database"

	"github.com/google/uuid"
)

type Todo struct {
	Title       string    `json:"title"`
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
}

func DatabaseListToList(dbTodos []database.GetTodoListRow) []Todo {
	todos := []Todo{}

	for _, dbTodo := range dbTodos {
		todos = append(todos, Todo(dbTodo))
	}
	return todos
}

func DatabaseTodoToTodo(dbTodo database.GetTodoByIdRow) Todo {
	return Todo{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: dbTodo.Description,
	}
}
