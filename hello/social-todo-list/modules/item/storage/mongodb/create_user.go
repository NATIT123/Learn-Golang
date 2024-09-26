package storagemongo

import (
	"context"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mongo *mongoStore) CreateUser(ctx context.Context, data *mmongodb.User) error {

	collection := mongo.Client.Database("testing").Collection(mmongodb.User{}.CollectionName())

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data.Id = primitive.NewObjectID()

	if _, err := collection.InsertOne(ctx, data); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
