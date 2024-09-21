package biz

import (
	"context"
	"main/modules/item/models"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context,cond map[string]interface{}) (*models.TodoItem,error)
	UpdateItem(ctx context.Context,cond map[string]interface{},dataUpdate *models.TodoItemUpdate) error
}

type updateItemBiz struct {
	store UpdateItemStorage
}

func NewUpdateItemBiz(store UpdateItemStorage) *updateItemBiz{
	return &updateItemBiz{store: store}
}

func (biz *updateItemBiz) UpdateItemById(ctx context.Context,id int,dataUpdate *models.TodoItemUpdate) error{
	

	data,err:=biz.store.GetItem(ctx,map[string]interface{}{"id":id})
		
	if err!=nil{
		return err
	}


	if data.Status !=nil && *data.Status == models.ItemStatusDeleted{
		return models.ErrItemDeleted
	}

	if err:=biz.store.UpdateItem(ctx,map[string]interface{}{"id":id},dataUpdate);err!=nil{
		return err
	}

	return nil
}