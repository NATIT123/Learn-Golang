package storage

import (
	"context"
	"main/common"
	"main/modules/item/models/postgreSQL"
)

func (sql *sqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {

	if err := sql.db.Table(models.TodoItem{}.TableName()).
		Where(cond).
		Updates(map[string]interface{}{
			"status": models.ItemStatusDeleted.String(),
		}).Error; err != nil {
		return common.ErrCannotCreateEntity(models.EntityName, err)
	}

	return nil
}
