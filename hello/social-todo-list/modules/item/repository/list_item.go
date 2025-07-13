package repository

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

type ItemLikeStorage interface {
	GetItemLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listItemRepo struct {
	store     ListItemStorage
	likeStore ItemLikeStorage
	requester common.Requester
}

func NewListItemRepo(store ListItemStorage, likeStore ItemLikeStorage, requester common.Requester) *listItemRepo {
	return &listItemRepo{store: store, likeStore: likeStore, requester: requester}
}

func (repo *listItemRepo) ListItem(ctx context.Context,
	filter *models.Filter,
	paging *common.Paging, moreKeys ...string) ([]models.TodoItem, error) {

	ctxStore := context.WithValue(ctx, common.CurrentUser, repo.requester)
	data, err := repo.store.ListItem(ctxStore, filter, paging, moreKeys...)

	if err != nil {
		return nil, common.ErrCannotGetEntity(models.EntityName, err)
	}

	if len(data) == 0 {
		return data, nil
	}

	ids := make([]int, len(data))

	for i := range ids {
		ids[i] = data[i].Id
	}

	likeUserMap, err := repo.likeStore.GetItemLikes(ctxStore, ids)

	if err != nil {

	}

	for i := range data {
		data[i].LikedCount = likeUserMap[data[i].Id]
		data[i].Mask()

	}

	return data, nil
}
