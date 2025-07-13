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

	for i := range morekeys {
		db = db.Preload(morekeys[i])
	}
	db = db.Preload("Owner")

	//Seeking Paging
	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			// return nil, common.ErrDB(err)
			return nil, common.NewCustomError(
				err,
				"Invalid cursor: unable to decode UID",
				"ErrInvalidCursor",
			)
		}
		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Select("*").Order("id desc").
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		result[len(result)-1].Mask()
		paging.NextCursor = result[len(result)-1].FakeId.String()
	} else {
		paging.NextCursor = ""
	}

	return result, nil
}
