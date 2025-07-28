package ginitem

import (
	"main/common"
	biz "main/modules/item/biz/postgreSQL"
	models "main/modules/item/models/postgreSQL"
	"main/modules/item/repository"
	storage "main/modules/item/storage/postgreSQL"
	"main/modules/item/storage/postgreSQL/restapi"
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItem(serviceCtx goservice.ServiceContext) func(*gin.Context) {
	return func(c *gin.Context) {

		db := serviceCtx.MustGet(common.PluginDBMain).(*gorm.DB)
		apiItemCaller := serviceCtx.MustGet(common.PluginAPIItem).(interface {
			GetServiceURL() string
		})

		var queryString struct {
			common.Paging
			models.Filter
		}

		if err := c.ShouldBind(&queryString); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryString.Paging.Process()

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := storage.NewSQLStore(db)

		///Repository
		serviceCtx := c.MustGet("serviceContext").(goservice.ServiceContext)
		likeStore := restapi.New(apiItemCaller.GetServiceURL(), serviceCtx.Logger("restapi.itemLikes"))
		repo := repository.NewListItemRepo(store, likeStore, requester)
		bussiness := biz.NewListItemBiz(repo, requester)

		result, err := bussiness.ListItem(c.Request.Context(), &queryString.Filter, &queryString.Paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &queryString.Filter, &queryString.Paging))
	}
}
