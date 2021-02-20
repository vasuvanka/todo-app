package database

import "github.com/vasuvanka/todo-app/backend/models"

//IDirectory - dir interface
type IDirectory interface {
	GetDirByID(id string) (models.Directory, error)
	GetDirsByQuery(query interface{}, skip, limit int) ([]models.Directory, error)
	GetUserDirs(userID, parentID string) ([]models.Directory, error)
	UpdateDir(models.Directory) error
	CreateDir(models.Directory) (models.Directory, error)
}
