package cmd

import (
	"fmt"
	"log"
	"main/common"
	"main/memcache"
	"main/middleware"
	storagemongo "main/modules/item/storage/mongodb"
	ginitemMongo "main/modules/item/transport/ginitem/mongodb"
	ginitem "main/modules/item/transport/ginitem/postgreSQL"
	ginuserlikeitem "main/modules/userlikeitem/transport/gin"
	"main/pubsub"
	"main/subscriber"

	"main/modules/upload"
	userStorage "main/modules/user/storage/postgreSQL"
	ginuser "main/modules/user/transport/ginuser/postgreSQL"
	"main/plugin/appredis"
	"main/plugin/rpccaller"
	"main/plugin/simple"
	"main/plugin/tokenprovider/jwt"
	"net/http"
	"os"
	"strings"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/plugin/storage/sdkgorm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("social-todo-list"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.PluginDBMain)),
		goservice.WithInitRunnable(jwt.NewTokenJWTProvider("main", common.PluginJWT)),
		goservice.WithInitRunnable(rpccaller.NewApiItemCaller(common.PluginAPIItem)),
		goservice.WithInitRunnable(pubsub.NewPubSub(common.PluginPubSub)),
		goservice.WithInitRunnable(appredis.NewRedisDB("redis", common.PluginRedis)),
		goservice.WithInitRunnable(simple.NewSimplePlugin("simple")),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start social TODO service",
	Run: func(cmd *cobra.Command, args []string) {

		service := newService()

		serviceLogger := service.Logger("service")
		API_KEY := os.Getenv("OpenWeatherMapApiKey")
		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		////MongoDb
		DB_MONGO := os.Getenv("DB_MONGO")
		DB_MONGO = strings.Replace(DB_MONGO, "db_username", os.Getenv("DB_MONGO_USER"), 1)
		DB_MONGO = strings.Replace(DB_MONGO, "<db_password>", os.Getenv("DB_MONGO_PASSWORD"), 1)
		store := storagemongo.CreateMongo(DB_MONGO)
		client := store.Client

		if err := client.Ping(nil, nil); err != nil {
			serviceLogger.Fatalln("Failed to connect to MongoDB:", err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recovery())

			type CanGetValue interface {
				GetValue() string
			}

			log.Println(service.MustGet("simple").(CanGetValue).GetValue())
			db := service.MustGet(common.PluginDBMain).(*gorm.DB)
			authStore := userStorage.NewSQLStore(db)
			authCache := memcache.NewUserCaching(memcache.NewRedisCache(service), authStore)

			middlewareAuth := middleware.RequiredAuth(authCache, service)

			engine.Static("/static", "./static")
			v1 := engine.Group("/v1")
			{
				v1.PUT("/upload", upload.Upload(service))
				users := v1.Group("/users")
				{
					users.POST("/register", ginuser.Register(service))
					users.POST("/login", ginuser.Login(service))
					users.GET("/profile", middlewareAuth, ginuser.Profile())
				}
				items := v1.Group("/items", middlewareAuth)
				{
					items.POST("", ginitem.CreateItem(service))
					items.GET("", ginitem.ListItem(service))
					items.GET("/:id", ginitem.GetItem(service))
					items.PATCH("/:id", ginitem.UpdateItem(service))
					items.DELETE("/:id", ginitem.DeleteItem(service))

					items.POST("/:id/like", ginuserlikeitem.UserLikeItem(service))
					items.DELETE("/:id/unlike", ginuserlikeitem.UserUnLikeItem(service))
					items.GET("/:id/user-liked-users", ginuserlikeitem.ListUserLiked(service))
				}

				rpc := v1.Group("rpc")
				{
					rpc.POST("/get_item_likes", ginuserlikeitem.GetItemLikes(service))
				}
			}

			v2 := engine.Group("/v2")
			{
				users := v2.Group("/users", middleware.Recovery())
				{
					users.POST("", ginitemMongo.CreateUser(client))
					users.GET("/:id", ginitemMongo.GetUser(client))
					users.PATCH("/:id", ginitemMongo.UpdateUser(client))
					users.DELETE("/:id", ginitemMongo.DeleteUser(client))
					users.GET("", ginitemMongo.ListUser(client))
				}
			}

			weather := engine.Group("/weather")
			{
				weather.GET("", ginitemMongo.GetWeather(API_KEY))
			}

			engine.GET("/ping", func(c *gin.Context) {
				go func() {
					defer common.Recovery()
					fmt.Println([]int{}[0])
				}()
				c.JSON(http.StatusOK, gin.H{
					"message": "Hello World",
				})
			})
		})

		_ = subscriber.NewEngine(service).Start()

		// subscriber.IncreaseLikeCountAfterUserLikeItem(service, context.Background())

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)

		}

	},
}

func Excucte() {
	rootCmd.AddCommand(outEnvCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}
