package repository

import (
	"context"

	"github.com/izzanzahrial/go-mongo-ex/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type student struct {
}

func (s *student) CreateStudent(ctx context.Context, collection *mongo.Collection, student *entity.Student) error {
	if _, err := collection.InsertOne(ctx, student); err != nil {
		return err
	}

	return nil
}

func (s *student) UpdateStudent(ctx context.Context, collection *mongo.Collection, student *entity.Student) error {
	bsonStudent, err := bson.Marshal(student)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": student.StudentId}
	update := bson.M{"$set": bsonStudent}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func (s *student) DeleteStudent(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (s *student) GetStudentById(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*entity.Student, error) {
	var student entity.Student

	if err := collection.FindOne(ctx, id).Decode(&student); err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *student) GetStudentByName(ctx context.Context, collection *mongo.Collection, name string) ([]*entity.Student, error) {
	filter := bson.M{"name": name}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []*entity.Student

	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	return students, nil
}

func (s *student) GetStudentsByClassOf(ctx context.Context, collection *mongo.Collection, classOf int) ([]*entity.Student, error) {
	filter := bson.M{"class_of": classOf}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []*entity.Student

	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	return students, nil
}

func (s *student) GetStudentsByClass(ctx context.Context, collection *mongo.Collection, classId primitive.ObjectID) ([]*entity.Student, error) {
	filter := bson.M{"class": classId}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []*entity.Student

	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	return students, nil
}

func (s *student) GetAllStudent(ctx context.Context, collection *mongo.Collection) ([]*entity.Student, error) {
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var students []*entity.Student

	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	return students, nil
}

func (s *student) DeleteAllStudent(ctx context.Context, collection *mongo.Collection) error {
	_, err := collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}

	return nil
}
