package user

import (
	"context"
	"os"

	"github.com/gustavocioccari/go-user-microservice/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	db *mongo.Client
}

type UserRepository interface {
	Create(user *models.User) (*mongo.InsertOneResult, error)
	FindById(id string) (*models.User, error)
}

const (
	collectionUsers = "users"
)

func NewUserRepository(db *mongo.Client) UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(user *models.User) (*mongo.InsertOneResult, error) {
	return r.db.Database(os.Getenv("DATABASE")).Collection(collectionUsers).InsertOne(context.TODO(), user)
}

func (r *repo) FindById(id string) (*models.User, error) {
	var result models.User

	r.db.Database(os.Getenv("DATABASE")).Collection(collectionUsers).FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)

	return &result, nil
}
