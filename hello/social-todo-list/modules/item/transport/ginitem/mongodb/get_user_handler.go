package ginitem

import (
	"main/common"
	bizmongo "main/modules/item/biz/mongodb"
	storagemongo "main/modules/item/storage/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(client *mongo.Client) func(*gin.Context) {
	return func(c *gin.Context) {

		id := c.Param("id")

		var objectID primitive.ObjectID

		if oid, err := primitive.ObjectIDFromHex(id); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		} else {
			objectID = oid
		}

		store := storagemongo.NewMongoStore(client)
		bussiness := bizmongo.NewGetUserBiz(store)

		data, err := bussiness.GetUserById(c.Request.Context(), objectID)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
