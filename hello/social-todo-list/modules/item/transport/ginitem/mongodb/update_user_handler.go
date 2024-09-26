package ginitem

import (
	"main/common"
	bizmongo "main/modules/item/biz/mongodb"
	mmongodb "main/modules/item/models/mongodb"
	storage "main/modules/item/storage/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(client *mongo.Client) func(*gin.Context) {
	return func(c *gin.Context) {

		var data mmongodb.UserUpdate

		id := c.Param("id")

		var objectID primitive.ObjectID

		if oid, err := primitive.ObjectIDFromHex(id); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		} else {
			objectID = oid
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return

		}

		store := storage.NewMongoStore(client)
		bussiness := bizmongo.NewUpdateUserBiz(store)

		if err := bussiness.UpdateUserById(c.Request.Context(), objectID, &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
