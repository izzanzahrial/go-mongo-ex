package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Teacher struct {
	TeacherId primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string               `bson:"name,omitempty" json:"name,omitempty"`
	Subjects  []string             `bson:"subjects,omitempty" json:"subjects,omitempty"`
	Class     []primitive.ObjectID `bson:"class,omitempty" json:"class,omitempty"`
}
