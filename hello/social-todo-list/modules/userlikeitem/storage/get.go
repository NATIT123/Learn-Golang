package storage

import (
	"context"
	"errors"
	"main/common"
	models "main/modules/userlikeitem/models"

	"gorm.io/gorm"
)

func (s *sqlStore) Get(ctx context.Context, userId int, itemId int) (*models.Like, error) {
	var data models.Like

	if err := s.db.Where("user_id = ? and item_id = ?", userId, itemId).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
