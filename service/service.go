package service

import (
	"context"

	"github.com/izzanzahrial/go-mongo-ex/db"
	"github.com/izzanzahrial/go-mongo-ex/entity"
	"github.com/izzanzahrial/go-mongo-ex/repository"
)

type Student interface {
	CreateStudent(ctx context.Context, student *entity.Student) error
	UpdateStudent(ctx context.Context, student *entity.Student) error
	DeleteStudent(ctx context.Context, id string) error
	GetStudentById(ctx context.Context, id string) (*entity.Student, error)
	GetStudentByName(ctx context.Context, name string) ([]*entity.Student, error)
	GetStudentsByClass(ctx context.Context, classId string) ([]*entity.Student, error)
	GetStudentsByClassOf(ctx context.Context, classOf int) ([]*entity.Student, error)
	GetAllStudent(ctx context.Context) ([]*entity.Student, error)
	DeleteAllStudent(ctx context.Context) error
}

func NewStudentService(coll string, repo repository.Student, con db.Connection) Student {
	return &student{
		collectionName: coll,
		repository:     repo,
		Connection:     con,
	}
}

type Teacher interface {
	CreateTeacher(ctx context.Context, teacher *entity.Teacher) error
	UpdateTeacher(ctx context.Context, teacher *entity.Teacher) error
	DeleteTeacher(ctx context.Context, id string) error
	GetTeacherById(ctx context.Context, id string) (*entity.Teacher, error)
	GetTeacherByName(ctx context.Context, name string) ([]*entity.Teacher, error)
	GetTeacherBySubject(ctx context.Context, subject string) ([]*entity.Teacher, error)
	GetTeacherByClass(ctx context.Context, classId string) ([]*entity.Teacher, error)
	GetAllTeacher(ctx context.Context) ([]*entity.Teacher, error)
	DeleteAllTeacher(ctx context.Context) error
}

func NewTeacherService(coll string, repo repository.Teacher, con db.Connection) Teacher {
	return &teacher{
		collectionName: coll,
		repository:     repo,
		Connection:     con,
	}
}

type Class interface {
	CreateClass(ctx context.Context, class *entity.Class) error
	UpdateClass(ctx context.Context, class *entity.Class) error
	DeleteClass(ctx context.Context, id string) error
	GetClassById(ctx context.Context, id string) (*entity.Class, error)
	GetClassByName(ctx context.Context, name string) ([]*entity.Class, error)
	GetClassPeriod(ctx context.Context, period string) ([]*entity.Class, error)
	GetAllClass(ctx context.Context) ([]*entity.Class, error)
	DeleteAllClass(ctx context.Context) error
}

func NewClassService(coll string, repo repository.Class, con db.Connection) Class {
	return &class{
		collectionName: coll,
		repository:     repo,
		Connection:     con,
	}
}
