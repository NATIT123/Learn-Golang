package routes

import (
	"be-ep/controllers"
	"be-ep/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userStorage models.UserStorage) {
	controllers := controllers.NewUserController(userStorage)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
}
