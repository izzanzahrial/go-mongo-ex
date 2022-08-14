package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	StudentId primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	ClassOf   int                `bson:"class_of" json:"class_of"`
}
