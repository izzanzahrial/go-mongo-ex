package service

import (
	"context"

	"github.com/izzanzahrial/go-mongo-ex/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student interface {
	CreateStudent(ctx context.Context, student *entity.Student) error
	UpdateStudent(ctx context.Context, student *entity.Student) error
	DeleteStudent(ctx context.Context, student *entity.Student) error
	GetStudentById(ctx context.Context, id primitive.ObjectID) (*entity.Student, error)
	GetStudentByName(ctx context.Context, name string) ([]*entity.Student, error)
	GetStudentsByClass(ctx context.Context, classId primitive.ObjectID) ([]*entity.Student, error)
	GetStudentsByClassOf(ctx context.Context, classOf int) ([]*entity.Student, error)
	GetAllStudent(ctx context.Context) ([]*entity.Student, error)
	DeleteAllStudent(ctx context.Context) error
}

type Teacher interface {
	CreateTeacher(ctx context.Context, teacher *entity.Teacher) error
	UpdateTeacher(ctx context.Context, teacher *entity.Teacher) error
	DeleteTeacher(ctx context.Context, teacher *entity.Teacher) error
	GetTeacherById(ctx context.Context, id primitive.ObjectID) (*entity.Teacher, error)
	GetTeacherByName(ctx context.Context, name string) ([]*entity.Teacher, error)
	GetTeacherBySubject(ctx context.Context, subject string) ([]*entity.Teacher, error)
	GetTeacherByClass(ctx context.Context, classId primitive.ObjectID) ([]*entity.Teacher, error)
	GetAllTeacher(ctx context.Context) ([]*entity.Teacher, error)
	DeleteAllTeacher(ctx context.Context) error
}

type Class interface {
	CreateClass(ctx context.Context, class *entity.Class) error
	UpdateClass(ctx context.Context, class *entity.Class) error
	DeleteClass(ctx context.Context, class *entity.Class) error
	GetClassById(ctx context.Context, id primitive.ObjectID) (*entity.Class, error)
	GetClassByName(ctx context.Context, name string) ([]*entity.Class, error)
	GetClassPeriod(ctx context.Context, period int) ([]*entity.Class, error)
	GetAllClass(ctx context.Context) ([]*entity.Class, error)
	DeleteAllClass(ctx context.Context) error
}
