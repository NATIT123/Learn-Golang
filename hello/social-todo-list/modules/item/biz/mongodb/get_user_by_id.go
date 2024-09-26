package bizmongo

import (
	"context"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	models "main/modules/item/models/postgreSQL"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*mmongodb.User, error)
}

type getUserBiz struct {
	store GetUserStorage
}

func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUserById(ctx context.Context, id primitive.ObjectID) (*mmongodb.User, error) {

	data, err := biz.store.GetUser(ctx, map[string]interface{}{"_id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(models.EntityName, err)
	}

	return data, nil
}
