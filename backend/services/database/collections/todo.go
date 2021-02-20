package collections

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/vasuvanka/todo-app/backend/models"
)

//GetTodosByQuery - get todos by query
func (db *Mongo) GetTodosByQuery(query interface{}, skip, limit int) ([]models.Todo, error) {
	var todos []models.Todo
	err := db.Todos().Find(query).Skip(skip).Limit(limit).All(&todos)
	return todos, err
}

//UpdateTodo - update todo
func (db *Mongo) UpdateTodo(todo models.Todo) error {
	return db.Todos().UpdateId(todo.ID, todo)
}

//GetTodoByID - get todo by ID
func (db *Mongo) GetTodoByID(id string) (models.Todo, error) {
	var todo models.Todo
	err := db.Todos().FindId(bson.ObjectIdHex(id)).One(&todo)
	return todo, err
}

//CreateTodo - create todo
func (db *Mongo) CreateTodo(todo models.Todo) (models.Todo, error) {
	todo.When = time.Now()
	todo.ID = bson.NewObjectId()
	todo.Removed = false
	err := db.Todos().Insert(todo)
	return todo, err
}

//GetUserTodos - get user todos
func (db *Mongo) GetUserTodos(id,dirID string, skip, limit int) ([]models.Todo, error) {
	var todos []models.Todo
	err := db.Todos().Find(bson.M{
		"createdBy": id,
		"removed":   false,
		"dirId": dirID,
	}).Skip(skip).Limit(limit).All(&todos)
	return todos, err
}
