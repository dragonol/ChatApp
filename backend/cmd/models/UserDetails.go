package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDetails struct {
	//Id primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Id          primitive.ObjectID
	Email       string
	Password    []byte
	PhoneNumber string
	DisplayName string
	DateOfBirth primitive.DateTime
}

func (u UserDetails) IsEmpty() bool {
	return u.Email == "" && len(u.Password) == 0
}
