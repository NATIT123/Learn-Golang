package main

import (
	"be-ep/configs"
	logic_pkg "be-ep/internal/logic"
	"be-ep/models"
	"be-ep/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func setHandler(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("value")

	if key == "" || val == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing key or value"})
		return
	}

	err := rdb.Set(ctx, key, val, 60*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key saved to Redis!"})
}

func getHandler(c *gin.Context) {
	key := c.Query("key")
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Key not found"})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"key": key, "value": val})
	}
}

func main() {
	fmt.Println(logic_pkg.SayHello("World"))

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("ENDPOINT"),
		Password: os.Getenv("PASSWORD"),
		Username: "default",
		DB:       0,
	})

	// Test Redis connection
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect Redis: %v", err)
	}

	configs.InitDB()
	userStorage := models.NewUserStorage(configs.DB)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Golang Gorm Server!",
			"status":  "success",
		})
	})
	routes.SetupRoutes(r, userStorage)
	r.GET("/set", setHandler)
	r.GET("/get", getHandler)
	log.Println("Server running on port 8888")
	r.Run(":8888")
}
