package ginitem

import (
	"main/common"
	bizmongo "main/modules/item/biz/mongodb"
	storage "main/modules/item/storage/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUser(client *mongo.Client) func(*gin.Context) {
	return func(c *gin.Context) {

		id := c.Param("id")

		var objectID primitive.ObjectID

		if oid, err := primitive.ObjectIDFromHex(id); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		} else {
			objectID = oid
		}

		store := storage.NewMongoStore(client)
		bussiness := bizmongo.NewDeleteUserBiz(store)

		if err := bussiness.DeleteUserById(c.Request.Context(), objectID); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
