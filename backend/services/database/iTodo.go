package database

import "github.com/vasuvanka/todo-app/backend/models"

//ITodo - todos interface
type ITodo interface {
	GetTodoByID(id string) (models.Todo, error)
	GetTodosByQuery(query interface{}, skip, limit int) ([]models.Todo, error)
	GetUserTodos(userID,dirID string, skip, limit int) ([]models.Todo, error)
	UpdateTodo(models.Todo) error
	CreateTodo(models.Todo) (models.Todo, error)
}
