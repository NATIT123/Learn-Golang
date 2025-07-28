package ginuserlikeitem

import (
	"main/common"
	storage "main/modules/item/storage/postgreSQL"
	"main/modules/userlikeitem/biz"
	"main/pubsub"
	"net/http"
	"strconv"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserUnLikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			// panic(common.ErrInvalidRequest(err))
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		ps := serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)
		store := storage.NewSQLStore(db)
		// itemStore := itemStorage.NewSQLStore(db)
		biz := biz.NewUserUnlikeItemBiz(store, ps)
		if err := biz.UnlikeItem(c.Request.Context(), requester.GetUserId(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
