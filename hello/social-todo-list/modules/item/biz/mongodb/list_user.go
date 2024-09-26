package bizmongo

import (
	"context"
	"main/common"
	mmongodb "main/modules/item/models/mongodb"
)

type ListUserStorage interface {
	ListUser(ctx context.Context,
		filter *mmongodb.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]mmongodb.User, error)
}

type listUserBiz struct {
	store ListUserStorage
}

func NewListUserBiz(store ListUserStorage) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUser(ctx context.Context,
	filter *mmongodb.Filter,
	paging *common.Paging) ([]mmongodb.User, error) {

	data, err := biz.store.ListUser(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotGetEntity(mmongodb.EntityName, err)
	}

	return data, nil
}
