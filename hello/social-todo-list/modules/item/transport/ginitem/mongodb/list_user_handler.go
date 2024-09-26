package ginitem

import (
	"main/common"
	bizmongo "main/modules/item/biz/mongodb"
	mmongodb "main/modules/item/models/mongodb"
	storagemongo "main/modules/item/storage/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ListUser(client *mongo.Client) func(*gin.Context) {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		paging.Process()

		var filter mmongodb.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storagemongo.NewMongoStore(client)
		bussiness := bizmongo.NewListUserBiz(store)

		result, err := bussiness.ListUser(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
