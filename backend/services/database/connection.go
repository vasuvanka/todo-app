package database

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/vasuvanka/todo-app/backend/services/database/collections"
)

//Db - database struct
type Db struct {
	collections.Mongo
}

//Connect - creates mongodb connection
func (db *Db) Connect(url string) error {
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	db.Mongo.Session = *session.Clone()
	log.Println("db connection established")
	return nil
}
