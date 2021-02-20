package collections

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/vasuvanka/todo-app/backend/models"
)

//GetUsersByQuery - get user by query
func (db *Mongo) GetUsersByQuery(query interface{}) ([]models.User, error) {
	var users []models.User
	err := db.Users().Find(query).All(&users)
	return users, err
}

//UpdateUser - update user
func (db *Mongo) UpdateUser(user models.User) error {
	return db.Users().UpdateId(user.ID, user)
}

//GetUserByID - get user by ID
func (db *Mongo) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := db.Users().FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

//GetUserByUsername - get user by email
func (db *Mongo) GetUserByUsername(email string) (models.User, error) {
	var user models.User
	err := db.Users().Find(bson.M{
		"email":   email,
		"removed": false,
	}).One(&user)
	return user, err
}

//CreateUser - create user
func (db *Mongo) CreateUser(user models.User) (models.User, error) {
	user.When = time.Now()
	user.Removed = false
	user.ID = bson.NewObjectId()
	err := db.Users().Insert(user)
	return user, err
}

//UserExist -  does user exist
func (db *Mongo) UserExist(email string) (bool, error) {
	count, err := db.Users().Find(bson.M{
		"email":   email,
		"removed": false,
	}).Count()
	return count > 0, err
}
