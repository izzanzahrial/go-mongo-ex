package main

import (
	"flag"
	"log"

	"github.com/izzanzahrial/go-mongo-ex/db"
	"github.com/izzanzahrial/go-mongo-ex/handler"
	"github.com/izzanzahrial/go-mongo-ex/repository"
	"github.com/izzanzahrial/go-mongo-ex/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	local       bool
	studentColl = "student"
	teacherColl = "teacher"
	classColl   = "class"
)

func init() {
	flag.BoolVar(&local, "local", true, "run service on local")
	flag.Parse()
}

func main() {
	if local {
		if err := godotenv.Load(); err != nil {
			log.Panic(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	studentRepo := repository.NewStudentRepository()
	teacherRepo := repository.NewTeacherRepository()
	classRepo := repository.NewClassRepository()

	studentService := service.NewStudentService(studentColl, studentRepo, conn)
	teacherService := service.NewTeacherService(teacherColl, teacherRepo, conn)
	classService := service.NewClassService(classColl, classRepo, conn)

	studentHandler := handler.NewStudentHandler(studentService)
	teacherHandler := handler.NewTeacherHandler(teacherService)
	classHandler := handler.NewClassHandler(classService)

	e := echo.New()

	s := e.Group("/api/v1/students")
	s.GET("", studentHandler.GetAllStudent)
	s.POST("", studentHandler.CreateStudent)
	s.PUT("", studentHandler.UpdateStudent)
	s.DELETE("", studentHandler.DeleteAllStudent)
	s.GET("/:id", studentHandler.GetStudentById)
	s.DELETE("/:id", studentHandler.DeleteStudent)
	s.GET("/name", studentHandler.GetStudentByName)
	s.GET("/class", studentHandler.GetStudentsByClass)
	s.GET("/class-of", studentHandler.GetStudentsByClassOf)

	t := e.Group("/api/v1/teachers")
	t.GET("", teacherHandler.GetAllTeacher)
	t.POST("", teacherHandler.CreateTeacher)
	t.PUT("", teacherHandler.UpdateTeacher)
	t.DELETE("", teacherHandler.DeleteAllTeacher)
	t.DELETE("/:id", teacherHandler.DeleteTeacher)
	t.GET("/:id", teacherHandler.GetTeacherById)
	t.GET("/name", teacherHandler.GetTeacherByName)
	t.GET("/subject", teacherHandler.GetTeacherBySubject)
	t.GET("/class", teacherHandler.GetTeacherByClass)

	c := e.Group("/api/v1/class")
	c.GET("", classHandler.GetAllClass)
	c.POST("", classHandler.CreateClass)
	c.PUT("", classHandler.UpdateClass)
	c.DELETE("", classHandler.DeleteAllClass)
	c.DELETE("/:id", classHandler.DeleteClass)
	c.GET("/:id", classHandler.GetClassById)
	c.GET("/name", classHandler.GetClassByName)
	c.GET("/period", classHandler.GetClassPeriod)

	e.Logger.Fatal(e.Start(":1323"))
}
