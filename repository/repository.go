package repository

import (
	"context"

	"github.com/izzanzahrial/go-mongo-ex/db"
	"github.com/izzanzahrial/go-mongo-ex/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentRepository interface {
	CreateStudent(ctx context.Context, conn db.Connection, student *entity.Student) error
	UpdateStudent(ctx context.Context, conn db.Connection, student *entity.Student) error
	DeleteStudent(ctx context.Context, conn db.Connection, student *entity.Student) error
	GetStudentById(ctx context.Context, conn db.Connection, id primitive.ObjectID) (*entity.Student, error)
	GetStudentByName(ctx context.Context, conn db.Connection, name string) (*entity.Student, error)
	GetStudentsByClassOf(ctx context.Context, conn db.Connection, classOf int) ([]*entity.Student, error)
	GetStudentsByClass(ctx context.Context, conn db.Connection, classId primitive.ObjectID) ([]*entity.Student, error)
	GetAllStudent(ctx context.Context, conn db.Connection) ([]*entity.Student, error)
	DeleteAllStudent(ctx context.Context, conn db.Connection) error
}

type TeacherRepository interface {
	CreateTeacher(ctx context.Context, conn db.Connection, teacher *entity.Teacher) error
	UpdateTeacher(ctx context.Context, conn db.Connection, teacher *entity.Teacher) error
	DeleteTeacher(ctx context.Context, conn db.Connection, teacher *entity.Teacher) error
	GetTeacherById(ctx context.Context, conn db.Connection, id primitive.ObjectID) (*entity.Teacher, error)
	GetTeacherByName(ctx context.Context, conn db.Connection, name string) (*entity.Teacher, error)
	GetTeacherBySubject(ctx context.Context, conn db.Connection, subject string) ([]*entity.Teacher, error)
	GetTeacherByClass(ctx context.Context, conn db.Connection, classId primitive.ObjectID) ([]*entity.Teacher, error)
	GetAllTeacher(ctx context.Context, conn db.Connection) ([]*entity.Teacher, error)
	DeleteAllTeacher(ctx context.Context, conn db.Connection) error
}

type ClassRepository interface {
	CreateClass(ctx context.Context, conn db.Connection, class *entity.Class) error
	UpdateClass(ctx context.Context, conn db.Connection, class *entity.Class) error
	DeleteClass(ctx context.Context, conn db.Connection, class *entity.Class) error
	GetClassById(ctx context.Context, conn db.Connection, id primitive.ObjectID) (*entity.Class, error)
	GetClassByName(ctx context.Context, conn db.Connection, name string) (*entity.Class, error)
	GetClassPeriod(ctx context.Context, conn db.Connection, period int) (*entity.Class, error)
	GetAllClass(ctx context.Context, conn db.Connection) ([]*entity.Class, error)
	DeleteAllClass(ctx context.Context, conn db.Connection) error
}
