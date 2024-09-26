package bizmongo

import (
	"context"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteUserStorage interface {
	GetUser(ctx context.Context, cond map[string]interface{}) (*mmongodb.User, error)
	DeleteUser(ctx context.Context, cond map[string]interface{}) error
}

type deleteUserBiz struct {
	store DeleteUserStorage
}

func NewDeleteUserBiz(store DeleteUserStorage) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUserById(ctx context.Context, id primitive.ObjectID) error {

	data, err := biz.store.GetUser(ctx, map[string]interface{}{"_id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(mmongodb.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(mmongodb.EntityName, err)
	}

	if !data.Active {
		return common.NewCustomError(mmongodb.ErrUserDeleted, "User has been deleted", "ErrUserDeleted")
	}

	if err := biz.store.DeleteUser(ctx, map[string]interface{}{"_id": id}); err != nil {
		return common.ErrCannotDeleteEntity(mmongodb.EntityName, err)
	}

	return nil
}
