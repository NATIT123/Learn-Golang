package biz

import (
	"context"
	"log"
	"main/common"
	models "main/modules/userlikeitem/models"
)

type UserUnlikeItemStore interface {
	Find(ctx context.Context, userId int, itemId int) (*models.Like, error)
	Delete(ctx context.Context, userId int, itemId int) error
}

type DecreaseItemStorage interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikeItemBiz struct {
	store     UserUnlikeItemStore
	itemStore DecreaseItemStorage
}

func NewUserUnlikeItemBiz(store UserUnlikeItemStore, itemStore DecreaseItemStorage) *userUnlikeItemBiz {
	return &userUnlikeItemBiz{store: store, itemStore: itemStore}
}

func (biz *userUnlikeItemBiz) UnlikeItem(ctx context.Context, userId int, itemId int) error {
	_, err := biz.store.Find(ctx, userId, itemId)

	// Delete if data existed
	if err == common.RecordNotFound {
		return models.ErrDidNotLikeItem(err)
	}

	if err != nil {
		return models.ErrCannotUnlikeItem(err)
	}

	if err := biz.store.Delete(ctx, userId, itemId); err != nil {
		return models.ErrCannotUnlikeItem(err)
	}

	go func() {
		defer common.Recovery()
		if err := biz.itemStore.DecreaseLikeCount(ctx, itemId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
