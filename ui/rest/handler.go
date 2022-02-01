package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
}

func NotFound(c echo.Context, data interface{}) error {
	err, match := data.(error)
	if match {
		data = echo.Map{"error": err.Error()}
	}

	return c.JSON(http.StatusNotFound, data)
}

func BadRequest(c echo.Context, data interface{}) error {
	err, match := data.(error)
	if match {
		data = echo.Map{"error": err.Error()}
	}

	return c.JSON(http.StatusBadRequest, data)
}
