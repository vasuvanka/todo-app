package collections

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/vasuvanka/todo-app/backend/models"
)

//GetDirsByQuery - get dirs by query
func (db *Mongo) GetDirsByQuery(query interface{}, skip, limit int) ([]models.Directory, error) {
	var dirs []models.Directory
	err := db.Dirs().Find(query).Skip(skip).Limit(limit).All(&dirs)
	return dirs, err
}

//UpdateDir - update directory
func (db *Mongo) UpdateDir(dir models.Directory) error {
	return db.Dirs().UpdateId(dir.ID, dir)
}

//GetDirByID - get dir by ID
func (db *Mongo) GetDirByID(id string) (models.Directory, error) {
	var dir models.Directory
	err := db.Dirs().FindId(bson.ObjectIdHex(id)).One(&dir)
	return dir, err
}

//CreateDir - create dir
func (db *Mongo) CreateDir(dir models.Directory) (models.Directory, error) {
	dir.When = time.Now()
	dir.ID = bson.NewObjectId()
	dir.Removed = false
	err := db.Dirs().Insert(dir)
	return dir, err
}

//GetUserDirs - get user dirs
func (db *Mongo) GetUserDirs(id, parentID string) ([]models.Directory, error) {
	var dirs []models.Directory
	err := db.Dirs().Find(bson.M{
		"createdBy": id,
		"removed":   false,
		"parentId":  parentID,
	}).All(&dirs)
	return dirs, err
}
