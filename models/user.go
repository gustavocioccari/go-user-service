package models

import (
	"time"

	"github.com/go-playground/validator"
)

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id"`
	Email     string    `json:"email" bson:"email" validate:"required"`
	Password  string    `json:"password,omitempty" bson:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"createdAt"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updatedAt"`
}

func (u *User) Validate() error {
	return validator.New().Struct(u)
}
