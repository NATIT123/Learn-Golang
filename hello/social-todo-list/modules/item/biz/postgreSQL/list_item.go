package biz

import (
	"context"
	"main/common"
	models "main/modules/item/models/postgreSQL"
)

type ListItemStorage interface {
	ListItem(ctx context.Context,
		filter *models.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]models.TodoItem, error)
}

type listItemBiz struct {
	store     ListItemStorage
	requester common.Requester
}

func NewListItemBiz(store ListItemStorage, requester common.Requester) *listItemBiz {
	return &listItemBiz{store: store, requester: requester}
}

func (biz *listItemBiz) ListItem(ctx context.Context,
	filter *models.Filter,
	paging *common.Paging) ([]models.TodoItem, error) {

	ctxStore := context.WithValue(ctx, common.CurrentUser, biz.requester)
	data, err := biz.store.ListItem(ctxStore, filter, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotGetEntity(models.EntityName, err)
	}

	return data, nil
}
