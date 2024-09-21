package biz

import (
	"context"
	"main/modules/item/models"
)

type GetItemStorage interface {
	GetItem(ctx context.Context,cond map[string]interface{}) (*models.TodoItem,error)
}

type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz{
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItemById(ctx context.Context,id int) (*models.TodoItem,error){
	

	data,err:=biz.store.GetItem(ctx,map[string]interface{}{"id":id})
		
	if err!=nil{
		return nil,err
	}

	return data,nil
}