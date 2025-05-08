package ginitem

import (
	"fmt"
	"main/common"
	biz "main/modules/item/biz/postgreSQL"
	models "main/modules/item/models/postgreSQL"
	storage "main/modules/item/storage/postgreSQL"
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

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSQLStore(db)
		bussiness := biz.NewListItemBiz(store, requester)

		result, err := bussiness.ListItem(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		for i := range result {
			result[i].Mask()
		}
		fmt.Println(result[0].FakeId)

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
