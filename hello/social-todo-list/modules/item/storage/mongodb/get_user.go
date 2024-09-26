package storagemongo

import (
	"context"
	"errors"
	"fmt"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (mongoClient *mongoStore) GetUser(ctx context.Context, cond map[string]interface{}) (*mmongodb.User, error) {

	var data mmongodb.User

	collection := mongoClient.Client.Database("testing").Collection(mmongodb.User{}.CollectionName())

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println(cond)

	err := collection.FindOne(ctx, cond).Decode(&data)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, common.RecordNotFound
	} else if err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
