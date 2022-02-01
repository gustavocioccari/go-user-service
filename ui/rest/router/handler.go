package router

import (
	userService "github.com/gustavocioccari/go-user-microservice/service/user"
	healthController "github.com/gustavocioccari/go-user-microservice/ui/rest/controllers/health"
	userController "github.com/gustavocioccari/go-user-microservice/ui/rest/controllers/user"
	"github.com/gustavocioccari/go-user-microservice/ui/rest/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter(userService userService.UserService) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middlewares.Cors())
	e.Use(middlewares.Logger())

	userController := userController.NewUserController(userService)
	healthController := healthController.NewHealthController()

	v1 := e.Group("user-service")
	{
		v1.GET("/health", healthController.Status)

		groupUser := v1.Group("/users")
		{
			groupUser.POST("", userController.Create)
		}
	}

	return e
}
