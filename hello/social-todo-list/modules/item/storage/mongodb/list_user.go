package storagemongo

import (
	"context"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (mongoClient *mongoStore) ListUser(ctx context.Context,
	filter *mmongodb.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]mmongodb.User, error) {

	var result []mmongodb.User

	collection := mongoClient.Client.Database("testing").Collection(mmongodb.User{}.CollectionName())

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var pipeline []bson.D

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "active", Value: true}}}}

	pipeline = append(pipeline, matchStage)

	pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "active", Value: true}}}})

	if f := filter; f != nil {
		if v := f.Active; v != "" {
			pipeline = append(pipeline, bson.D{{Key: "$match", Value: bson.D{{Key: "active", Value: v}}}})
		}
	}

	sortStage := bson.D{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}}
	limitStage := bson.D{{Key: "$limit", Value: paging.Limit}}
	skipStage := bson.D{{Key: "$skip", Value: ((paging.Page - 1) * paging.Limit)}}

	pipeline = append(pipeline, sortStage, limitStage, skipStage)

	cursor, err := collection.Aggregate(ctx, mongo.Pipeline(pipeline))
	if err != nil {
		return nil, common.ErrDB(err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, common.ErrDB(err)
	} else {
		paging.Total = int64(len(result))
	}

	return result, nil
}
