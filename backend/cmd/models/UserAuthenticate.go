package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAuthenticate struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password []byte             `bson:"password"`
}

func (u UserAuthenticate) IsEmpty() bool {
	return u.Email == "" && len(u.Password) == 0
}
