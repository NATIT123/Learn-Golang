package bizmongo

import (
	"context"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	models "main/modules/item/models/postgreSQL"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*mmongodb.User, error)
	UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *mmongodb.UserUpdate) error
}

type updateUserBiz struct {
	store UpdateUserStorage
}

func NewUpdateUserBiz(store UpdateUserStorage) *updateUserBiz {
	return &updateUserBiz{store: store}
}

func (biz *updateUserBiz) UpdateUserById(ctx context.Context, id primitive.ObjectID, dataUpdate *mmongodb.UserUpdate) error {

	data, err := biz.store.GetUser(ctx, map[string]interface{}{"_id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(models.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(models.EntityName, err)
	}

	if !data.Active {
		return common.ErrEntityDeleted(models.EntityName, mmongodb.ErrUserDeleted)
	}

	if err := biz.store.UpdateUser(ctx, map[string]interface{}{"_id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(models.EntityName, err)
	}

	return nil
}
