package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppID = primitive.ObjectID

func NewAppID() AppID {
	return primitive.NewObjectID()
}
