package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Class struct {
	Id     primitive.ObjectID `bson:"_id" json:"_id"`
	Name   string             `bson:"name" json:"name"`
	Period int                `bson:"period" json:"period"`
}
