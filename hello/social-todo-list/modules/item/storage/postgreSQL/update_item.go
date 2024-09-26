package storage

import (
	"context"
	"main/common"
	models "main/modules/item/models/postgreSQL"
)

func (sql *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *models.TodoItemUpdate) error {

	if err := sql.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return common.ErrCannotUpdateEntity(models.EntityName, err)
	}

	return nil
}
