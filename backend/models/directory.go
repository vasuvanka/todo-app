package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Directory - directroy structure
type Directory struct {
	Title     string        `json:"title" bson:"title"`
	Removed   bool          `json:"removed" bson:"removed"`
	CreatedBy string        `json:"-" bson:"createdBy"`
	ID        bson.ObjectId `json:"id" bson:"_id"`
	When      time.Time     `json:"when,omitempty" bson:"when,omitempty"`
	ParentID  string        `json:"parentId" bson:"parentId"`
}
