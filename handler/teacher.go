package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/izzanzahrial/go-mongo-ex/entity"
	"github.com/izzanzahrial/go-mongo-ex/service"
	"github.com/labstack/echo/v4"
)

type teacher struct {
	service service.Teacher
}

func (t *teacher) CreateTeacher(c echo.Context) error {
	var teacher entity.Teacher
	teacher.Name = c.FormValue("name")
	teacher.Subjects = []string{c.FormValue("subject")}

	if err := t.service.CreateTeacher(context.Background(), &teacher); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (t *teacher) UpdateTeacher(c echo.Context) error {
	var teacher entity.Teacher

	if err := json.NewDecoder(c.Request().Body).Decode(&teacher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	defer c.Request().Body.Close()

	if err := t.service.UpdateTeacher(context.Background(), &teacher); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, nil)
}

func (t *teacher) DeleteTeacher(c echo.Context) error {
	id := c.Param("id")

	if err := t.service.DeleteTeacher(context.Background(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (t *teacher) GetTeacherById(c echo.Context) error {
	id := c.Param("id")

	teacher, err := t.service.GetTeacherById(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, teacher)
}

func (t *teacher) GetTeacherByName(c echo.Context) error {
	name := c.QueryParam("name")

	teachers, err := t.service.GetTeacherByName(context.Background(), name)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, teachers)
}

func (t *teacher) GetTeacherBySubject(c echo.Context) error {
	subject := c.QueryParam("subject")

	teachers, err := t.service.GetTeacherBySubject(context.Background(), subject)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, teachers)
}

func (t *teacher) GetTeacherByClass(c echo.Context) error {
	classID := c.QueryParam("class_id")

	teachers, err := t.service.GetTeacherByClass(context.Background(), classID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, teachers)
}

func (t *teacher) GetAllTeacher(c echo.Context) error {
	teachers, err := t.service.GetAllTeacher(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, teachers)
}

func (t *teacher) DeleteAllTeacher(c echo.Context) error {
	if err := t.service.DeleteAllTeacher(context.Background()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
