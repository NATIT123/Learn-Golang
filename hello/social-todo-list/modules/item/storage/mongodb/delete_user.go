package storagemongo

import (
	"context"
	"errors"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (mongoClient *mongoStore) DeleteUser(ctx context.Context, cond map[string]interface{}) error {

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

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "active", Value: false}}}}

	if _, err := collection.UpdateOne(ctx, cond, update); err != nil {
		return common.ErrCannotUpdateEntity(mmongodb.EntityName, err)
	}

	return nil
}
