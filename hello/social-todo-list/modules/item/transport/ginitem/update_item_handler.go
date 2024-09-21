package ginitem

import (
	"main/common"
	"main/modules/item/biz"
	"main/modules/item/models"
	"main/modules/item/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		
		var data models.TodoItemUpdate

		id,err:=strconv.Atoi(c.Param("id"))

		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return 
		}

		if err:=c.ShouldBind(&data);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		store:= storage.NewSQLStore(db)
		bussiness :=biz.NewUpdateItemBiz(store)

		if err:= bussiness.UpdateItemById(c.Request.Context(),id,&data);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,common.SimpleSuccessResponse(true))
	}
}