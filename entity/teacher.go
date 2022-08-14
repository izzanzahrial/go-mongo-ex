package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Teacher struct {
	TeacherId primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `bson:"name" json:"name"`
	Age       int                `bson:"age" json:"age"`
	Subjects  []string           `bson:"subjects" json:"subjects"`
}
