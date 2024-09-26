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

	fmt.Println(dataUpdate)

	if _, err := collection.UpdateOne(ctx, cond, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(mmongodb.EntityName, err)
	}

	return nil
}
