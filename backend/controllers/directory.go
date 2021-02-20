package controllers

import (
	"errors"

	"github.com/vasuvanka/todo-app/backend/models"
	"github.com/vasuvanka/todo-app/backend/services/database"
)

var iDir database.IDirectory

//NewDirectoryController - will get dir interface
func NewDirectoryController(dbDir database.IDirectory) {
	iDir = dbDir
}

//GetUserDirs - get users dir
func GetUserDirs(userID, parentID string) ([]models.Directory, error) {
	return iDir.GetUserDirs(userID, parentID)
}

//CreateDir - create dir
func CreateDir(dir models.Directory) (models.Directory, error) {
	dir.Removed = false
	return iDir.CreateDir(dir)
}

//UpdateDir - update dir
func UpdateDir(dir models.Directory) error {
	dbDir, err := iDir.GetDirByID(dir.ID.Hex())
	if err != nil {
		return err
	}
	if dir.CreatedBy != dbDir.CreatedBy {
		return errors.New("todo was created by others")
	}
	dbDir.Title = dir.Title
	return iDir.UpdateDir(dir)
}

//GetDirByID - get dir by id
func GetDirByID(id string) (models.Directory, error) {
	return iDir.GetDirByID(id)
}

//DeleteDir - delete dir
func DeleteDir(id string) error {
	dir, err := iDir.GetDirByID(id)
	if err != nil {
		return err
	}
	dir.Removed = true
	return iDir.UpdateDir(dir)
}
