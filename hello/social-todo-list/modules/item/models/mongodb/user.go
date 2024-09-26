package mmongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	EntityName = "User"
)

var (
	ErrTitleIsBlank = errors.New("name can not be blank")
	ErrItemDeleted  = errors.New("user is deleted")
)

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Photo    string             `json:"photo" bson:"photo"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"`
	Active   bool               `json:"active" bson:"active"`
}

func (User) CollectionName() string { return "users" }

type UserUpdate struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Photo    string `json:"photo" bson:"photo"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
	Active   bool   `json:"active" bson:"active"`
}
