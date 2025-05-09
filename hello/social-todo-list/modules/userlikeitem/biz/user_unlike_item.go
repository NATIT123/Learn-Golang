package biz

import (
	"context"
	"main/common"
	models "main/modules/userlikeitem/models"
)

type UserUnlikeItemStore interface {
	Find(ctx context.Context, userId int, itemId int) (*models.Like, error)
	Delete(ctx context.Context, userId int, itemId int) error
}

type userUnlikeItemBiz struct {
	store UserUnlikeItemStore
}

func NewUserUnlikeItemBiz(store UserUnlikeItemStore) *userUnlikeItemBiz {
	return &userUnlikeItemBiz{store: store}
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

	return nil
}
