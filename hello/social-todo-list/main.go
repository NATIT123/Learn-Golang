package main

import (
	"fmt"
	"log"
	"main/common"
	"main/middleware"
	storagemongo "main/modules/item/storage/mongodb"
	storage "main/modules/item/storage/postgreSQL"
	ginitemMongo "main/modules/item/transport/ginitem/mongodb"
	ginitem "main/modules/item/transport/ginitem/postgreSQL"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	DB_CONN_STR := os.Getenv("DB_CONN_STR")

	db := storage.CreateSQL(DB_CONN_STR)

	////MongoDb
	DB_MONGO := os.Getenv("DB_MONGO")
	store := storagemongo.CreateMongo(DB_MONGO)
	client := store.Client

	// now:=time.Now().UTC()

	// item := TodoItem{
	// 	Id:	1,
	// 	Title:"This is item 1",
	// 	Description:"This is item 1",
	// 	Status:	ItemStatusDoing,
	// 	CreatedAt: &now,
	// 	UpdatedAt: nil,
	// }

	// ///Convert JSONData =byte[],err
	// jsonData,err:= json.Marshal(item)

	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }

	// ///Convert to JSON
	// fmt.Println(string(jsonData))

	// var item2 TodoItem

	// if err:=json.Unmarshal([]byte(jsonData),&item2); err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(item2)

	/////For
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Printf("Sum: %d\n", sum)

	///While in Go
	sum = 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println("Sum:", sum)

	var pow = []int{1, 2, 3, 4, 5}
	for i, value := range pow {
		fmt.Printf("Index: %d,Value: %d\n", i, value)
	}

	///CRUD:Create,Read,Update,Delete

	r := gin.Default()

	// r.Use(middleware.Recovery())

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items", middleware.Recovery())
		{
			items.POST("", ginitem.CreateItem(db))
			items.GET("", ginitem.ListItem(db))
			items.GET("/:id", ginitem.GetItem(db))
			items.PATCH("/:id", ginitem.UpdateItem(db))
			items.DELETE("/:id", ginitem.DeleteItem(db))
		}
	}

	v2 := r.Group("/v2")
	{
		users := v2.Group("/users")
		{
			users.POST("", ginitemMongo.CreateUser(client))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(PORT) // listen and serve on 0.0.0.0:3000(for windows "localhost:3000")

}
