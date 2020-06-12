package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDetails struct {
	Id          primitive.ObjectID `bson:"_id"`
	Email       string
	DisplayName string
}
