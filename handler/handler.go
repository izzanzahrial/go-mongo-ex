package handler

import (
	"github.com/izzanzahrial/go-mongo-ex/service"
	"github.com/labstack/echo/v4"
)

type Student interface {
	CreateStudent(c echo.Context) error
	UpdateStudent(c echo.Context) error
	DeleteStudent(c echo.Context) error
	GetStudentById(c echo.Context) error
	GetStudentByName(c echo.Context) error
	GetStudentsByClass(c echo.Context) error
	GetStudentsByClassOf(c echo.Context) error
	GetAllStudent(c echo.Context) error
	DeleteAllStudent(c echo.Context) error
}

func NewStudentHandler(service service.Student) Student {
	return &student{service: service}
}

type Teacher interface {
	CreateTeacher(c echo.Context) error
	UpdateTeacher(c echo.Context) error
	DeleteTeacher(c echo.Context) error
	GetTeacherById(c echo.Context) error
	GetTeacherByName(c echo.Context) error
	GetTeacherBySubject(c echo.Context) error
	GetTeacherByClass(c echo.Context) error
	GetAllTeacher(c echo.Context) error
	DeleteAllTeacher(c echo.Context) error
}

func NewTeacherHandler(service service.Teacher) Teacher {
	return &teacher{service: service}
}

type Class interface {
	CreateClass(c echo.Context) error
	UpdateClass(c echo.Context) error
	DeleteClass(c echo.Context) error
	GetClassById(c echo.Context) error
	GetClassByName(c echo.Context) error
	GetClassPeriod(c echo.Context) error
	GetAllClass(c echo.Context) error
	DeleteAllClass(c echo.Context) error
}

func NewClassHandler(service service.Class) Class {
	return &class{service: service}
}
