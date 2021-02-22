package controllers

import (
	"errors"

	"github.com/vasuvanka/todo-app/backend/models"
	"github.com/vasuvanka/todo-app/backend/services/database"
	"github.com/vasuvanka/todo-app/backend/shared"
)

var iTodo database.ITodo

//NewTodoController - will get user interface
func NewTodoController(dbTodo database.ITodo) {
	iTodo = dbTodo
}

//GetUserTodos - get user todos
func GetUserTodos(userID,dirID string, skip, limit int) ([]models.Todo, error) {
	return iTodo.GetUserTodos(userID, dirID, skip, limit)
}

//GetTodoByID - get todo by id
func GetTodoByID(id string) (models.Todo, error) {
	return iTodo.GetTodoByID(id)
}

// UpdateTodo - update todo
func UpdateTodo(todo models.Todo) (models.Todo, error) {
	dbTodo, err := iTodo.GetTodoByID(todo.ID.Hex())
	if err != nil {
		return models.Todo{}, err
	}
	if todo.CreatedBy != dbTodo.CreatedBy {
		return models.Todo{}, errors.New("todo was created by others")
	}
	dbTodo.Title = todo.Title
	dbTodo.Description = todo.Description
	dbTodo.Priority = todo.Priority
	dbTodo.Status = todo.Status
	err = iTodo.UpdateTodo(dbTodo)
	if err != nil {
		return models.Todo{}, err
	}
	return dbTodo, nil
}

// DeleteTodoByID - delete todod
func DeleteTodoByID(ID, userID string) error {
	todo, err := GetTodoByID(ID)
	if err != nil {
		return err
	}
	if todo.CreatedBy != userID {
		return errors.New("todo was created by others")
	}
	todo.Removed = true
	return iTodo.UpdateTodo(todo)
}

//CreateTodo - create todo
func CreateTodo(todo models.Todo) (models.Todo, error) {
	todo.Status = shared.StatusCreated
	todo.Type = shared.TypeSelf
	return iTodo.CreateTodo(todo)
}

//ShareTodo - share todo with other user
func ShareTodo(todoID, fromUserID, toUserEmail string) error {
	dbTodo, err := iTodo.GetTodoByID(todoID)
	if err != nil {
		return err
	}
	if dbTodo.CreatedBy != fromUserID {
		return errors.New("todo was created by others")
	}

	user, err := iUser.GetUserByID(fromUserID)
	if err != nil {
		return err
	}

	if user.Email == toUserEmail {
		return errors.New("operation not permitted")
	}

	exist, err := iUser.UserExist(toUserEmail)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("user does not exist")
	}

	toUser, err := iUser.GetUserByUsername(toUserEmail)
	if err != nil {
		return err
	}
	dbTodo.Type = shared.TypeShared
	dbTodo.SharedBy = dbTodo.CreatedBy
	dbTodo.ParentID = dbTodo.ID.Hex()
	dbTodo.Status = shared.StatusCreated
	dbTodo.CreatedBy = toUser.ID.Hex()
	dbTodo.Removed = false
	dbTodo.DirID = "0"
	_, err = iTodo.CreateTodo(dbTodo)
	return err
}
