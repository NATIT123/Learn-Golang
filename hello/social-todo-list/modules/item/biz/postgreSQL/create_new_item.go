package biz

import (
	"context"
	"main/common"
	"main/modules/item/models"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *models.TodoItemCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *models.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return common.NewCustomError(models.ErrTitleIsBlank, "field not found", "ErrFieldNotFound")
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(models.EntityName, err)
	}

	return nil
}
