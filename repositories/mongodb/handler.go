package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDB() (*mongo.Client, error) {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin123@localhost:27017/admin"))
	if err != nil {
		return nil, err
	}

	if err := db.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	db.Database(os.Getenv("DATABASE")).Collection("users").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true).SetName("email"),
	})

	return db, nil
}
