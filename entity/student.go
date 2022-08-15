package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	StudentId primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string               `bson:"name,omitempty" json:"name,omitempty"`
	ClassOf   int                  `bson:"class_of,omitempty" json:"class_of,omitempty"`
	Class     []primitive.ObjectID `bson:"class,omitempty" json:"class,omitempty"`
}
