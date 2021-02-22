package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Todo Database Object
type Todo struct {
	Title       string        `vd:"len($)>0; msg:'Title required'" json:"title" bson:"name"`
	Type        string        `json:"type" bson:"type"` // self/shared
	Description string        `vd:"len($)>3; msg:'Description required and should be greater than 3 char'" json:"description" bson:"description"`
	Removed     bool          `json:"removed" bson:"removed"`
	Priority    string        `json:"priority" bson:"priority"`
	Status      string        `json:"status" bson:"status"`
	CreatedBy   string        `json:"-" bson:"createdBy"`
	SharedBy    string        `json:"-" bson:"sharedBy"`
	ID          bson.ObjectId `json:"id" bson:"_id"`
	When        time.Time     `json:"when,omitempty" bson:"when,omitempty"`
	DueDate     time.Time     `json:"dueDate,omitempty" bson:"dueDate,omitempty"`
	ParentID    string 		  `json:"-" bson:"parentId,omitempty"`
	DirID    	string        `vd:"len($)>0; msg:'Directory Id required'" json:"dirId" bson:"dirId"`
}
