package models

import "github.com/4nthem/State/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	User struct {
		Id    bson.ObjectId `json:"id" bson:"_id"`
		Name  string        `json:"name" bson:"name"`
		Email string        `json:"email" bson:"email"`
	}
)
