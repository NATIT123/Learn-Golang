package storage

import (
	"context"
	"main/common"
	models "main/modules/item/models/postgreSQL"

	"gorm.io/gorm"
)

func (sql *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *models.TodoItemUpdate) error {

	if err := sql.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return common.ErrCannotUpdateEntity(models.EntityName, err)
	}

	return nil
}

func (s *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(models.TodoItem{}.TableName()).Where("query = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(models.TodoItem{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
