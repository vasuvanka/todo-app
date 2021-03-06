package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Directory - directroy structure
type Directory struct {
	Title     string        `vd:"len($)>0; msg:'Title required'" json:"title" bson:"title"`
	Removed   bool          `json:"removed" bson:"removed"`
	CreatedBy string        `json:"-" bson:"createdBy"`
	ID        bson.ObjectId `json:"id" bson:"_id"`
	When      time.Time     `json:"when,omitempty" bson:"when,omitempty"`
	ParentID  string        `vd:"len($)>0; msg:'ParentId required'" json:"parentId" bson:"parentId"`
}
