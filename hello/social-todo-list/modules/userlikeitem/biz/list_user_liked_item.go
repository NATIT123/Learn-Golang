package biz

import (
	"context"
	"main/common"
	models "main/modules/userlikeitem/models"
)

type ListUserLikeItemStore interface {
	ListUsers(
		ctx context.Context,
		itemId int,
		paging *common.Paging,
	) ([]common.SimpleUser, error)
}

type listUserLikeItemBiz struct {
	store ListUserLikeItemStore
}

func NewListUserLikeItemBiz(store ListUserLikeItemStore) *listUserLikeItemBiz {
	return &listUserLikeItemBiz{store: store}
}

func (biz *listUserLikeItemBiz) ListLikedItem(
	ctx context.Context,
	itemId int,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	result, err := biz.store.ListUsers(ctx, itemId, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(models.EntityName, err)
	}

	return result, nil
}
