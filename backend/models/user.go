package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//User - User Profile
type User struct {
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	Role       string        `json:"role" bson:"role"`
	Password   string        `json:"-" bson:"password"`
	ID         bson.ObjectId `json:"id" bson:"_id"`
	When       time.Time     `json:"when" bson:"when"`
	ResetToken string        `json:"-" bson:"resetToken,omitempty"`
	Removed    bool          `json:"removed" bson:"removed"`
}
