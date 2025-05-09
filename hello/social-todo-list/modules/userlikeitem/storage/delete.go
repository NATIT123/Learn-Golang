package storage

import (
	"context"
	"main/common"
	models "main/modules/userlikeitem/models"
)

func (s *sqlStore) Delete(ctx context.Context, userId int, itemId int) error {
	var data models.Like

	if err := s.db.Table(data.TableName()).Where("user_id = ? and item_id = ?", userId, itemId).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
