package storage

import (
	"context"
	"main/common"
	models "main/modules/userlikeitem/models"
)

func (s *sqlStore) Create(ctx context.Context, data *models.Like) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
