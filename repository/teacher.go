package repository

import (
	"context"

	"github.com/izzanzahrial/go-mongo-ex/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type teacherRepository struct {
	collection *mongo.Collection
}

func (tr *teacherRepository) CreateTeacher(ctx context.Context, collection *mongo.Collection, teacher *entity.Teacher) error {
	if _, err := collection.InsertOne(ctx, teacher); err != nil {
		return err
	}

	return nil
}

func (tr *teacherRepository) UpdateTeacher(ctx context.Context, collection *mongo.Collection, teacher *entity.Teacher) error {
	bsonTeacher, err := bson.Marshal(teacher)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": teacher.TeacherId}
	update := bson.M{"$set": bsonTeacher}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func (tr *teacherRepository) DeleteTeacher(ctx context.Context, collection *mongo.Collection, teacher *entity.Teacher) error {
	filter := bson.M{"_id": teacher.TeacherId}
	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (tr *teacherRepository) GetTeacherById(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*entity.Teacher, error) {
	var teacher entity.Teacher

	if err := collection.FindOne(ctx, id).Decode(&teacher); err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (tr *teacherRepository) GetTeacherByName(ctx context.Context, collection *mongo.Collection, name string) ([]*entity.Teacher, error) {
	filter := bson.M{"name": name}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var teachers []*entity.Teacher

	if err := cursor.All(ctx, &teachers); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (tr *teacherRepository) GetTeacherBySubject(ctx context.Context, collection *mongo.Collection, subject string) ([]*entity.Teacher, error) {
	filter := bson.M{"subject": subject}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var teachers []*entity.Teacher

	if err := cursor.All(ctx, &teachers); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (tr *teacherRepository) GetTeacherByClass(ctx context.Context, collection *mongo.Collection, classId primitive.ObjectID) ([]*entity.Teacher, error) {
	filter := bson.M{"class": classId}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var teachers []*entity.Teacher

	if err := cursor.All(ctx, &teachers); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (tr *teacherRepository) GetAllTeacher(ctx context.Context, collection *mongo.Collection) ([]*entity.Teacher, error) {
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var teachers []*entity.Teacher

	if err := cursor.All(ctx, &teachers); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (tr *teacherRepository) DeleteAllTeacher(ctx context.Context, collection *mongo.Collection) error {
	_, err := collection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}

	return nil
}
