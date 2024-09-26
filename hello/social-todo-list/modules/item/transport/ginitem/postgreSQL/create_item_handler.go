package ginitem

import (
	"main/common"
	"main/modules/item/biz/postgreSQL"
	"main/modules/item/models/postgreSQL"
	"main/modules/item/storage/postgreSQL"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.TodoItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)
		bussiness := biz.NewCreateItemBiz(store)

		if err := bussiness.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
