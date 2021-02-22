package collections

import (
	"github.com/globalsign/mgo"
	"github.com/vasuvanka/todo-app/backend/shared"
)

//Mongo - mongo struct
type Mongo struct {
	Session mgo.Session
}

//Users - Users collection
func (db *Mongo) Users() *mgo.Collection {
	return db.Session.DB(shared.Database).C(shared.Users)
}

//Todos - Todos collection
func (db *Mongo) Todos() *mgo.Collection {
	return db.Session.DB(shared.Database).C(shared.Todos)
}

//Dirs - directory collection
func (db *Mongo) Dirs() *mgo.Collection {
	return db.Session.DB(shared.Database).C(shared.Dirs)
}

//Index - will index db fields
func (db *Mongo) Index() error {
	// if err := db.Dirs().DropAllIndexes(); err != nil {
	// 	return err
	// }
	idx := mgo.Index{
		Key:      []string{"email"},
		Unique:   true,
		DropDups: true,
		Name:     "user_email_index",
	}
	// user indexes
	if err := db.Users().EnsureIndex(idx); err != nil {
		return err
	}

	idx = mgo.Index{
		Key:  []string{"removed", "name"},
		Name: "user_name_index",
	}
	// user indexes
	if err := db.Users().EnsureIndex(idx); err != nil {
		return err
	}

	idx = mgo.Index{
		Key:  []string{"removed", "title", "description", "createdBy"},
		Name: "todos_index",
	}
	// todo indexes
	if err := db.Todos().EnsureIndex(idx); err != nil {
		return err
	}

	idx = mgo.Index{
		Key:  []string{"removed", "title", "createdBy"},
		Name: "directory_index",
	}
	// directory indexes
	return db.Dirs().EnsureIndex(idx)
}
