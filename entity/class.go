package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Class struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Period   int                `bson:"period,omitempty" json:"period,omitempty"`
	Students []Student          `bson:"students,omitempty" json:"students,omitempty"`
}
