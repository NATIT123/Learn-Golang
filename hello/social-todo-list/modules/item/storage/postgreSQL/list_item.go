package storage

import (
	"context"
	"main/common"
	models "main/modules/item/models/postgreSQL"
)

func (sql *sqlStore) ListItem(ctx context.Context,
	filter *models.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]models.TodoItem, error) {

	var result []models.TodoItem

	db := sql.db.Where("status <> ?", "Delete")

	requester := ctx.Value(common.CurrentUser).(common.Requester)

	//Get items of requester only
	db = db.Where("user_id = ?", requester.GetUserId())

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(models.TodoItem{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	db = db.Preload("Owner")

	if err := db.Select("*").Order("id desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
