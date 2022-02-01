package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthController struct{}

type HealthController interface {
	Status(c echo.Context) error
}

func NewHealthController() HealthController {
	return &healthController{}
}

func (hc healthController) Status(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "Server is up and running"})
}
