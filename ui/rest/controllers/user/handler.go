package user

import (
	"log"
	"net/http"

	"github.com/gustavocioccari/go-user-microservice/models"
	userService "github.com/gustavocioccari/go-user-microservice/service/user"
	"github.com/gustavocioccari/go-user-microservice/ui/rest"
	"github.com/labstack/echo/v4"
)

type controller struct {
	userService userService.UserService
}

type UserController interface {
	Create(c echo.Context) error
}

func NewUserController(userService userService.UserService) UserController {
	return &controller{
		userService: userService,
	}
}

func (c controller) Create(ctx echo.Context) error {
	var user *models.User

	err := ctx.Bind(&user)
	if err != nil {
		return rest.InternalServerError(ctx, err)
	}
	log.Println("User binded", user.Email)

	if err = user.Validate(); err != nil {
		return rest.BadRequest(ctx, err)
	}
	log.Println("User validated", user.Email)

	userCreated, err := c.userService.Create(user)
	if err != nil {
		return rest.InternalServerError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, userCreated)
}
