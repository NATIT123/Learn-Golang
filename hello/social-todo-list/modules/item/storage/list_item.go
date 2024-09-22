package storage

import (
	"context"
	"main/common"
	"main/modules/item/models"
)

func (sql *sqlStore) ListItem(ctx context.Context,
	filter *models.Filter,
	paging *common.Paging,
	morekeys ...string,
	) ([]models.TodoItem,error){
	
		var result [] models.TodoItem


		db := sql.db.Where("status <> ?","Delete")


		if f:=filter;f!=nil{
			if v:=f.Status;v!=""{
				db = db.Where("status = ?",v)
			}
		}

		if	err:=db.Table(models.TodoItem{}.TableName()).
		Count(&paging.Total).Error;err!=nil{
			return nil,err
		}

		if err:=db.Order("id desc").
		Offset((paging.Page-1)*paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error;err!=nil{
			return nil,err
		}

	return result,nil
}