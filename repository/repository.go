package repository

import (
	"context"

	"github.com/izzanzahrial/go-mongo-ex/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Student interface {
	CreateStudent(ctx context.Context, collection *mongo.Collection, student *entity.Student) error
	UpdateStudent(ctx context.Context, collection *mongo.Collection, student *entity.Student) error
	DeleteStudent(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) error
	GetStudentById(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*entity.Student, error)
	GetStudentByName(ctx context.Context, collection *mongo.Collection, name string) ([]*entity.Student, error)
	GetStudentsByClass(ctx context.Context, collection *mongo.Collection, classId primitive.ObjectID) ([]*entity.Student, error)
	GetStudentsByClassOf(ctx context.Context, collection *mongo.Collection, classOf int) ([]*entity.Student, error)
	GetAllStudent(ctx context.Context, collection *mongo.Collection) ([]*entity.Student, error)
	DeleteAllStudent(ctx context.Context, collection *mongo.Collection) error
}

type Teacher interface {
	CreateTeacher(ctx context.Context, collection *mongo.Collection, teacher *entity.Teacher) error
	UpdateTeacher(ctx context.Context, collection *mongo.Collection, teacher *entity.Teacher) error
	DeleteTeacher(ctx context.Context, collection *mongo.Collection, teacher *entity.Teacher) error
	GetTeacherById(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*entity.Teacher, error)
	GetTeacherByName(ctx context.Context, collection *mongo.Collection, name string) ([]*entity.Teacher, error)
	GetTeacherBySubject(ctx context.Context, collection *mongo.Collection, subject string) ([]*entity.Teacher, error)
	GetTeacherByClass(ctx context.Context, collection *mongo.Collection, classId primitive.ObjectID) ([]*entity.Teacher, error)
	GetAllTeacher(ctx context.Context, collection *mongo.Collection) ([]*entity.Teacher, error)
	DeleteAllTeacher(ctx context.Context, collection *mongo.Collection) error
}

type Class interface {
	CreateClass(ctx context.Context, collection *mongo.Collection, class *entity.Class) error
	UpdateClass(ctx context.Context, collection *mongo.Collection, class *entity.Class) error
	DeleteClass(ctx context.Context, collection *mongo.Collection, class *entity.Class) error
	GetClassById(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*entity.Class, error)
	GetClassByName(ctx context.Context, collection *mongo.Collection, name string) ([]*entity.Class, error)
	GetClassPeriod(ctx context.Context, collection *mongo.Collection, period int) ([]*entity.Class, error)
	GetAllClass(ctx context.Context, collection *mongo.Collection) ([]*entity.Class, error)
	DeleteAllClass(ctx context.Context, collection *mongo.Collection) error
}
