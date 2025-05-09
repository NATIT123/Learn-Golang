package storage

import (
	"context"
	"gorm.io/gorm"
	model "main/modules/userlikeitem/models"
)

type sqlStore struct {
	db *gorm.DB
}

// Find implements biz.UserUnlikeItemStore.
func (s *sqlStore) Find(ctx context.Context, userId int, itemId int) (*model.Like, error) {
	panic("unimplemented")
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
