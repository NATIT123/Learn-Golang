package cmd

import (
	"fmt"
	"main/common"
	"main/component/tokenprovider/jwt"
	"main/middleware"
	storagemongo "main/modules/item/storage/mongodb"
	ginitemMongo "main/modules/item/transport/ginitem/mongodb"
	ginitem "main/modules/item/transport/ginitem/postgreSQL"
	"main/modules/upload"
	userStorage "main/modules/user/storage/postgreSQL"
	ginuser "main/modules/user/transport/ginuser/postgreSQL"
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
			db := service.MustGet(common.PluginDBMain).(*gorm.DB)
			authStore := userStorage.NewSQLStore(db)
			tokenprovider := jwt.NewTokenJWTProvider("jwt", os.Getenv("JWT_SECRET_KEY"))
			middlewareAuth := middleware.RequiredAuth(authStore, tokenprovider)

			engine.Static("/static", "./static")
			v1 := engine.Group("/v1")
			{
				v1.PUT("/upload", upload.Upload(db))
				users := v1.Group("/users")
				{
					users.POST("/register", ginuser.Register(db))
					users.POST("/login", ginuser.Login(db, tokenprovider))
					users.GET("/profile", middlewareAuth, ginuser.Profile())
				}
				items := v1.Group("/items", middlewareAuth)
				{
					items.POST("", ginitem.CreateItem(db))
					items.GET("", ginitem.ListItem(db))
					items.GET("/:id", ginitem.GetItem(db))
					items.PATCH("/:id", ginitem.UpdateItem(db))
					items.DELETE("/:id", ginitem.DeleteItem(db))
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
