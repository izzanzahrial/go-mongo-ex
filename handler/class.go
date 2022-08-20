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

type class struct {
	service service.Class
}

func (cl *class) CreateClass(c echo.Context) error {
	period, err := strconv.Atoi(c.FormValue("period"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var class entity.Class
	class.Name = c.FormValue("name")
	class.Period = period

	if err := cl.service.CreateClass(context.Background(), &class); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (cl *class) UpdateClass(c echo.Context) error {
	var class entity.Class

	if err := json.NewDecoder(c.Request().Body).Decode(&class); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	defer c.Request().Body.Close()

	if err := cl.service.UpdateClass(context.Background(), &class); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, nil)
}

func (cl *class) DeleteClass(c echo.Context) error {
	id := c.Param("id")

	if err := cl.service.DeleteClass(context.Background(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}

func (cl *class) GetClassById(c echo.Context) error {
	id := c.Param("id")

	class, err := cl.service.GetClassById(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, class)
}

func (cl *class) GetClassByName(c echo.Context) error {
	name := c.QueryParam("name")

	classes, err := cl.service.GetClassByName(context.Background(), name)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, classes)
}

func (cl *class) GetClassPeriod(c echo.Context) error {
	period := c.QueryParam("name")

	classes, err := cl.service.GetClassPeriod(context.Background(), period)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, classes)
}

func (cl *class) GetAllClass(c echo.Context) error {
	classes, err := cl.service.GetAllClass(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusFound, classes)
}

func (cl *class) DeleteAllClass(c echo.Context) error {
	if err := cl.service.DeleteAllClass(context.Background()); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, nil)
}
