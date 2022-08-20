package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/izzanzahrial/go-mongo-ex/entity"
	"github.com/izzanzahrial/go-mongo-ex/service"
	"github.com/labstack/echo/v4"
)

type student struct {
	service service.Student
}

func (s *student) CreateStudent(c echo.Context) error {
	classOf, err := strconv.Atoi(c.FormValue("class_of"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var student entity.Student
	student.Name = c.FormValue("name")
	student.ClassOf = classOf

	if err := s.service.CreateStudent(context.Background(), &student); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (s *student) UpdateStudent(c echo.Context) error {
	var student entity.Student

	if err := json.NewDecoder(c.Request().Body).Decode(&student); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	defer c.Request().Body.Close()

	if err := s.service.UpdateStudent(context.Background(), &student); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, nil)
}

func (s *student) DeleteStudent(c echo.Context) error {
	id := c.Param("id")

	if err := s.service.DeleteStudent(context.Background(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (s *student) GetStudentById(c echo.Context) error {
	id := c.Param("id")

	student, err := s.service.GetStudentById(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, student)
}

func (s *student) GetStudentByName(c echo.Context) error {
	name := c.QueryParam("name")

	students, err := s.service.GetStudentByName(context.Background(), name)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, students)
}

func (s *student) GetStudentsByClass(c echo.Context) error {
	classID := c.QueryParam("class_id")

	students, err := s.service.GetStudentsByClass(context.Background(), classID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, students)
}

func (s *student) GetStudentsByClassOf(c echo.Context) error {
	classOf, err := strconv.Atoi(c.QueryParam("class_of"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	students, err := s.service.GetStudentsByClassOf(context.Background(), classOf)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, students)
}

func (s *student) GetAllStudent(c echo.Context) error {
	students, err := s.service.GetAllStudent(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, students)
}

func (s *student) DeleteAllStudent(c echo.Context) error {
	if err := s.service.DeleteAllStudent(context.Background()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
