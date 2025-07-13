package ginuserlikeitem

import (
	"main/common"
	itemStorage "main/modules/item/storage/postgreSQL"
	"main/modules/userlikeitem/biz"
	model "main/modules/userlikeitem/models"
	storage "main/modules/userlikeitem/storage"
	"net/http"
	"strconv"
	"time"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserLikeItem(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			// panic(common.ErrInvalidRequest(err))
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		store := storage.NewSQLStore(db)
		itemStore := itemStorage.NewSQLStore(db)
		biz := biz.NewUserLikeItemBiz(store, itemStore)
		now := time.Now().UTC()
		if err := biz.LikeItem(c.Request.Context(), &model.Like{
			UserId:    requester.GetUserId(),
			ItemId:    id,
			CreatedAt: &now,
		}); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
