package service

import (
	"context"

	"github.com/izzanzahrial/go-mongo-ex/db"
	"github.com/izzanzahrial/go-mongo-ex/entity"
	"github.com/izzanzahrial/go-mongo-ex/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type teacher struct {
	collectionName string
	repository     repository.Teacher
	db.Connection
}

func (t *teacher) CreateTeacher(ctx context.Context, teacher *entity.Teacher) error {
	collection := t.DB().Collection(t.collectionName)

	if err := t.repository.CreateTeacher(ctx, collection, teacher); err != nil {
		return err
	}

	return nil
}

func (t *teacher) UpdateTeacher(ctx context.Context, teacher *entity.Teacher) error {
	collection := t.DB().Collection(t.collectionName)

	if err := t.repository.UpdateTeacher(ctx, collection, teacher); err != nil {
		return err
	}

	return nil
}

func (t *teacher) DeleteTeacher(ctx context.Context, teacher *entity.Teacher) error {
	collection := t.DB().Collection(t.collectionName)

	if err := t.repository.DeleteTeacher(ctx, collection, teacher); err != nil {
		return err
	}

	return nil
}

func (t *teacher) GetTeacherById(ctx context.Context, id primitive.ObjectID) (*entity.Teacher, error) {
	collection := t.DB().Collection(t.collectionName)

	teacher, err := t.repository.GetTeacherById(ctx, collection, id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (t *teacher) GetTeacherByName(ctx context.Context, name string) ([]*entity.Teacher, error) {
	collection := t.DB().Collection(t.collectionName)

	teachers, err := t.repository.GetTeacherByName(ctx, collection, name)
	if err != nil {
		return nil, err
	}

	return teachers, nil
}

func (t *teacher) GetTeacherBySubject(ctx context.Context, subject string) ([]*entity.Teacher, error) {
	collection := t.DB().Collection(t.collectionName)

	teachers, err := t.repository.GetTeacherBySubject(ctx, collection, subject)
	if err != nil {
		return nil, err
	}

	return teachers, nil
}

func (t *teacher) GetTeacherByClass(ctx context.Context, classId primitive.ObjectID) ([]*entity.Teacher, error) {
	collection := t.DB().Collection(t.collectionName)

	teachers, err := t.repository.GetTeacherByClass(ctx, collection, classId)
	if err != nil {
		return nil, err
	}

	return teachers, nil
}

func (t *teacher) GetAllTeacher(ctx context.Context) ([]*entity.Teacher, error) {
	collection := t.DB().Collection(t.collectionName)

	teachers, err := t.repository.GetAllTeacher(ctx, collection)
	if err != nil {
		return nil, err
	}

	return teachers, nil
}

func (t *teacher) DeleteAllTeacher(ctx context.Context) error {
	collection := t.DB().Collection(t.collectionName)

	if err := t.repository.DeleteAllTeacher(ctx, collection); err != nil {
		return err
	}

	return nil
}
