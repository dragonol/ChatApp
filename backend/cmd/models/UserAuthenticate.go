package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAuthenticate struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email    string
	Password []byte
}

func (u UserAuthenticate) IsEmpty() bool {
	return u.Email == "" && len(u.Password) == 0
}
