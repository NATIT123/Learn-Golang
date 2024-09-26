package mmongodb

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	Id       bson.ObjectID `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Photo    string        `json:"photo" bson:"photo"`
	Password string        `json:"password" bson:"password"`
	Role     string        `json:"role" bson:"role"`
	Active   bool          `json:"active" bson:"active"`
}

func (User) CollectionName() string { return "users" }
