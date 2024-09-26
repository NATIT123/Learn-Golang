package biz

import (
	"context"
	"main/common"
	"main/modules/item/models/postgreSQL"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*models.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *models.TodoItemUpdate) error
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{store: store}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *models.TodoItemUpdate) error {

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(models.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(models.EntityName, err)
	}

	if data.Status != nil && *data.Status == models.ItemStatusDeleted {
		return common.ErrEntityDeleted(models.EntityName, models.ErrItemDeleted)
	}

	if err := biz.store.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(models.EntityName, err)
	}

	return nil
}
