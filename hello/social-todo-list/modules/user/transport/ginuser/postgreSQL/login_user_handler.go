package ginuser

import (
	"main/common"
	"main/modules/user/biz"
	models "main/modules/user/models/postgreSQL"
	storage "main/modules/user/storage/postgreSQL"
	"main/plugin/tokenprovider"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, tokenProvider tokenprovider.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData models.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		bcrypt := common.NewBcryptHash()
		expiry := 60 * 60 * 24 * 30 // 30 days

		business := biz.NewLoginBusiness(store, tokenProvider, bcrypt, expiry)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
