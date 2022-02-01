package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_custom}] ${method} ${status} ${host} ${path} ${latency_human}\n",
		CustomTimeFormat: "02/01/2006 15:04:05",
	})
}
