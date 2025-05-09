package biz

import (
	"context"
	models "main/modules/userlikeitem/models"
)

type UserLikeItemStore interface {
	Create(ctx context.Context, data *models.Like) error
}

type userLikeItemBiz struct {
	store UserLikeItemStore
}

func NewUserLikeItemBiz(store UserLikeItemStore) *userLikeItemBiz {
	return &userLikeItemBiz{store: store}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *models.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return models.ErrCannotLikeItem(err)
	}
	return nil
}
