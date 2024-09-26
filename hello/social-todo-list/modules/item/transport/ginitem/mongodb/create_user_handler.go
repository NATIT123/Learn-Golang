package ginitem

import (
	"main/common"
	biz "main/modules/item/biz/mongodb"
	mmongodb "main/modules/item/models/mongodb"
	storage "main/modules/item/storage/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(client *mongo.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		var data mmongodb.User

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewMongoStore(client)
		bussiness := biz.NewCreateUserBiz(store)

		if err := bussiness.CreateNewUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
