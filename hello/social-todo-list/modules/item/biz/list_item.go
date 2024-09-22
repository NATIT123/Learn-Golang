package biz

import (
	"context"
	"main/common"
	"main/modules/item/models"
)

type ListItemStorage interface {
	ListItem(ctx context.Context,
		filter *models.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]models.TodoItem, error)
}

type listItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *listItemBiz {
	return &listItemBiz{store: store}
}

func (biz *listItemBiz) ListItem(ctx context.Context,
	filter *models.Filter,
	paging *common.Paging) ([]models.TodoItem, error) {

	data, err := biz.store.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotGetEntity(models.EntityName, err)
	}

	return data, nil
}
