package bizmongo

import (
	"context"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
	models "main/modules/item/models/postgreSQL"
	"strings"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *mmongodb.User) error
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *mmongodb.User) error {
	name := strings.TrimSpace(data.Name)

	if name == "" {
		return common.NewCustomError(models.ErrTitleIsBlank, "field not found", "ErrFieldNotFound")
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(models.EntityName, err)
	}

	return nil
}
