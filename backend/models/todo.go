package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Todo Database Object
type Todo struct {
	Title       string        `json:"title" bson:"name"`
	Type        string        `json:"type" bson:"type"` // self/shared
	Description string        `json:"description" bson:"description"`
	Removed     bool          `json:"removed" bson:"removed"`
	Priority    string        `json:"priority" bson:"priority"`
	Status      string        `json:"status" bson:"status"`
	CreatedBy   string        `json:"-" bson:"createdBy"`
	SharedBy    string        `json:"-" bson:"sharedBy"`
	ID          bson.ObjectId `json:"id" bson:"_id"`
	When        time.Time     `json:"when,omitempty" bson:"when,omitempty"`
	DueDate     time.Time     `json:"dueDate,omitempty" bson:"dueDate,omitempty"`
	ParentID    string 		  `json:"-" bson:"parentId,omitempty"`
	DirID    	string        `json:"dirId" bson:"dirId"`
}
