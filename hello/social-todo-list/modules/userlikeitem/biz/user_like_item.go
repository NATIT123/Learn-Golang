package biz

import (
	"context"
	"log"
	"main/common"
	models "main/modules/userlikeitem/models"
)

type UserLikeItemStore interface {
	Create(ctx context.Context, data *models.Like) error
}

type userLikeItemBiz struct {
	store     UserLikeItemStore
	itemStore IncreaseItemStorage
}

type IncreaseItemStorage interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

func NewUserLikeItemBiz(store UserLikeItemStore, itemStore IncreaseItemStorage) *userLikeItemBiz {
	return &userLikeItemBiz{store: store, itemStore: itemStore}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *models.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return models.ErrCannotLikeItem(err)
	}

	go func() {
		defer common.Recovery()
		if err := biz.itemStore.IncreaseLikeCount(ctx, data.ItemId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
