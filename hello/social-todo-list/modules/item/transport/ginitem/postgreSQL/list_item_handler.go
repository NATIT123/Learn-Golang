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

func ListItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		paging.Process()

		var filter models.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)
		bussiness := biz.NewListItemBiz(store)

		result, err := bussiness.ListItem(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
