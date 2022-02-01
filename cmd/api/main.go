package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gustavocioccari/go-user-microservice/repositories/mongodb"
	"github.com/gustavocioccari/go-user-microservice/repositories/mongodb/user"
	userService "github.com/gustavocioccari/go-user-microservice/service/user"
	"github.com/gustavocioccari/go-user-microservice/ui/rest/router"
)

func main() {
	db, err := mongodb.GetDB()
	if err != nil {
		log.Println(err)
	}
	userRepository := user.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)

	router := router.SetupRouter(userService)

	if err := router.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatalln("Error on start rest:", err)
	}
}
