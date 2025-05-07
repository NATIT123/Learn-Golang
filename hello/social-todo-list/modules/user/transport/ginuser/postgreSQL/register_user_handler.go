package ginuser

import (
	"main/common"
	"main/modules/user/biz"
	models "main/modules/user/models/postgreSQL"
	storage "main/modules/user/storage/postgreSQL"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			// errRes := common.NewRestErr(http.StatusBadRequest, "invalid request", err)
			// c.JSON(http.StatusBadRequest, errRes)
			// return
			panic(err)
		}

		store := storage.NewSQLStore(db)
		bcrypt := common.NewBcryptHash()
		biz := biz.NewRegisterBusiness(store, bcrypt)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			// errRes := common.NewRestErr(http.StatusBadRequest, err.Error(), err)
			// c.JSON(http.StatusBadRequest, errRes)
			// return
			panic(err)
		}

		// data.Mask(common.DbTypeUser)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
