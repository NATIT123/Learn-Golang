package biz

import (
	"context"
	"main/modules/item/models"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context,data *models.TodoItemCreation) error
}

type createItemBiz struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz{
	return &createItemBiz{store: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context,data *models.TodoItemCreation) error{
	title :=strings.TrimSpace(data.Title)


	if title == ""{
		return models.ErrTitleIsBlank
	}

	if err:=biz.store.CreateItem(ctx,data);err!=nil{
		return err
	}

	return nil
}