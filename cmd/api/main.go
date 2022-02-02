package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gustavocioccari/go-user-microservice/repositories/mongodb"
	"github.com/gustavocioccari/go-user-microservice/repositories/mongodb/user"
	"github.com/gustavocioccari/go-user-microservice/service/kafka"
	userService "github.com/gustavocioccari/go-user-microservice/service/user"
	"github.com/gustavocioccari/go-user-microservice/ui/rest/router"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, err := mongodb.GetDB()
	if err != nil {
		log.Println(err)
	}

	userRepository := user.NewUserRepository(db)
	kafkaService := kafka.NewKafkaService()
	userService := userService.NewUserService(userRepository, kafkaService)

	router := router.SetupRouter(userService)

	if err := router.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatalln("Error on start rest:", err)
	}
}
