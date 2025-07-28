package biz

import (
	"context"
	"log"
	"main/common"
	models "main/modules/userlikeitem/models"
	"main/pubsub"
)

type UserLikeItemStore interface {
	Create(ctx context.Context, data *models.Like) error
}

type userLikeItemBiz struct {
	store UserLikeItemStore
	// itemStore IncreaseItemStorage
	ps pubsub.PubSub
}

// type IncreaseItemStorage interface {
// 	IncreaseLikeCount(ctx context.Context, id int) error
// }

func NewUserLikeItemBiz(store UserLikeItemStore,

	// itemStore IncreaseItemStorage,

	ps pubsub.PubSub) *userLikeItemBiz {
	return &userLikeItemBiz{store: store,
		//  itemStore: itemStore,
		ps: ps}
}

func (biz *userLikeItemBiz) LikeItem(ctx context.Context, data *models.Like) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return models.ErrCannotLikeItem(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserLikedItem, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	// go func() {
	// 	defer common.Recovery()
	// 	if err := biz.itemStore.IncreaseLikeCount(ctx, data.ItemId); err != nil {
	// 		log.Println(err)
	// 	}
	// }()

	// job := asyncjob.NewJob(func(ctx context.Context) error {
	// 	if err := biz.itemStore.IncreaseLikeCount(ctx, data.ItemId); err != nil {
	// 		return err
	// 	}

	// 	return nil
	// })

	// if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
	// 	log.Println(err)
	// }

	return nil
}
