package storagemongo

import (
	"context"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (mongo *mongoStore) CreateUser(ctx context.Context, data *mmongodb.User) error {

	collection := mongo.Client.Database("test").Collection(mmongodb.User{}.CollectionName())

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := collection.InsertOne(ctx, bson.D{{Key: "name", Value: "Tu"}}); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
