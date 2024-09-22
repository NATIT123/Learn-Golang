package storage

import (
	"context"
	"main/common"
	"main/modules/item/models"

	"gorm.io/gorm"
)

func (sql *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*models.TodoItem, error) {
	var data models.TodoItem

	if err := sql.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
