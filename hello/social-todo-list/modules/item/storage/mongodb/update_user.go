package storagemongo

import (
	"context"
	"errors"
	"fmt"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (mongoClient *mongoStore) UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *mmongodb.UserUpdate) error {

	var data mmongodb.User

	collection := mongoClient.Client.Database("testing").Collection(mmongodb.User{}.CollectionName())

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, cond).Decode(&data)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return common.RecordNotFound
	} else if err != nil {
		return common.ErrDB(err)
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: dataUpdate.Name},
		{Key: "email", Value: dataUpdate.Email},
		{Key: "photo", Value: dataUpdate.Photo},
		{Key: "password", Value: dataUpdate.Password},
		{Key: "role", Value: dataUpdate.Role},
		{Key: "active", Value: dataUpdate.Active}}}}

	fmt.Println(update)

	if _, err := collection.UpdateOne(ctx, cond, update); err != nil {
		return common.ErrCannotUpdateEntity(mmongodb.EntityName, err)
	}

	return nil
}
