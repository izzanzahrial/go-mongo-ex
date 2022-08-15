package service

import (
	"context"

	"github.com/izzanzahrial/go-mongo-ex/db"
	"github.com/izzanzahrial/go-mongo-ex/entity"
	"github.com/izzanzahrial/go-mongo-ex/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type student struct {
	collectionName string
	repository     repository.Student
	db.Connection
}

func (s *student) CreateStudent(ctx context.Context, student *entity.Student) error {
	collection := s.DB().Collection(s.collectionName)

	if err := s.repository.CreateStudent(ctx, collection, student); err != nil {
		return err
	}

	return nil
}

func (s *student) UpdateStudent(ctx context.Context, student *entity.Student) error {
	collection := s.DB().Collection(s.collectionName)

	if err := s.repository.UpdateStudent(ctx, collection, student); err != nil {
		return err
	}

	return nil
}

func (s *student) DeleteStudent(ctx context.Context, id string) error {
	collection := s.DB().Collection(s.collectionName)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err := s.repository.DeleteStudent(ctx, collection, objectID); err != nil {
		return err
	}

	return nil
}

func (s *student) GetStudentById(ctx context.Context, id string) (*entity.Student, error) {
	collection := s.DB().Collection(s.collectionName)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	student, err := s.repository.GetStudentById(ctx, collection, objectID)
	if err != nil {
		return nil, err
	}

	return student, err
}

func (s *student) GetStudentByName(ctx context.Context, name string) ([]*entity.Student, error) {
	collection := s.DB().Collection(s.collectionName)

	students, err := s.repository.GetStudentByName(ctx, collection, name)
	if err != nil {
		return nil, err
	}

	return students, err
}

func (s *student) GetStudentsByClass(ctx context.Context, classId string) ([]*entity.Student, error) {
	collection := s.DB().Collection(s.collectionName)
	objectID, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		return nil, err
	}

	students, err := s.repository.GetStudentsByClass(ctx, collection, objectID)
	if err != nil {
		return nil, err
	}

	return students, err
}

func (s *student) GetStudentsByClassOf(ctx context.Context, classOf int) ([]*entity.Student, error) {
	collection := s.DB().Collection(s.collectionName)

	students, err := s.repository.GetStudentsByClassOf(ctx, collection, classOf)
	if err != nil {
		return nil, err
	}

	return students, err
}

func (s *student) GetAllStudent(ctx context.Context) ([]*entity.Student, error) {
	collection := s.DB().Collection(s.collectionName)

	students, err := s.repository.GetAllStudent(ctx, collection)
	if err != nil {
		return nil, err
	}

	return students, err
}

func (s *student) DeleteAllStudent(ctx context.Context) error {
	collection := s.DB().Collection(s.collectionName)

	err := s.repository.DeleteAllStudent(ctx, collection)
	if err != nil {
		return err
	}

	return err
}
