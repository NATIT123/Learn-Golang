package ginuser

import (
	"main/common"
	"main/modules/user/biz"
	models "main/modules/user/models/postgreSQL"
	storage "main/modules/user/storage/postgreSQL"
	"main/plugin/tokenprovider"
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(serviceCtx goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData models.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		store := storage.NewSQLStore(db)
		bcrypt := common.NewBcryptHash()
		expiry := 60 * 60 * 24 * 30 // 30 days
		tokenProvider := serviceCtx.MustGet(common.PluginJWT).(tokenprovider.Provider)
		business := biz.NewLoginBusiness(store, tokenProvider, bcrypt, expiry)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
