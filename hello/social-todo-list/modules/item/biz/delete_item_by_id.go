package biz

import (
	"context"
	"main/modules/item/models"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context,cond map[string]interface{}) (*models.TodoItem,error)
	DeleteItem(ctx context.Context,cond map[string]interface{}) error
}

type deleteItemBiz struct {
	store DeleteItemStorage
}

func NewDeleteItemBiz(store DeleteItemStorage) *deleteItemBiz{
	return &deleteItemBiz{store: store}
}

func (biz *deleteItemBiz) DeletetemById(ctx context.Context,id int) error{
	

	data,err:=biz.store.GetItem(ctx,map[string]interface{}{"id":id})
		
	if err!=nil{
		return err
	}


	if data.Status !=nil && *data.Status == models.ItemStatusDeleted{
		return models.ErrItemDeleted
	}

	if err:=biz.store.DeleteItem(ctx,map[string]interface{}{"id":id});err!=nil{
		return err
	}

	return nil
}