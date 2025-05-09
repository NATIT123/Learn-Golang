package ginuserlikeitem

import (
	"main/common"
	"main/modules/userlikeitem/biz"
	storage "main/modules/userlikeitem/storage"
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
		store := storage.NewSQLStore(db)
		biz := biz.NewUserUnlikeItemBiz(store)
		if err := biz.UnlikeItem(c.Request.Context(), requester.GetUserId(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
