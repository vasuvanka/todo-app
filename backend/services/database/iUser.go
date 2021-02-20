package database

import (
	"github.com/vasuvanka/todo-app/backend/models"
)

//IUser - user interface
type IUser interface {
	GetUsersByQuery(interface{}) ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	UpdateUser(models.User) error
	CreateUser(models.User) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	UserExist(username string) (bool, error)
}
