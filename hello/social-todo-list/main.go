package main

import (
	"fmt"
	"log"
	"main/common"
	"main/modules/item/models"
	"main/modules/item/transport/ginitem"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


  




func main(){

	err:=godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	  }
	PORT:=os.Getenv("PORT")
	DB_CONN_STR:=os.Getenv("DB_CONN_STR")


	// https://github.com/jackc/pgx
	///Connect PostgreSQL 
	dsn := DB_CONN_STR
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!=nil{
		log.Fatalln(err)
	}

	fmt.Println(db)

	

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
	sum:=0
	for i:=0; i<10 ; i++ {
		sum+=i
	}

	fmt.Printf("Sum: %d\n",sum)


	///While in Go
	sum= 1
	for sum<1000{
		sum+=sum
	}

	fmt.Println("Sum:",sum)


	var pow =[]int{1,2,3,4,5}
	for i,value :=range pow{
		fmt.Printf("Index: %d,Value: %d\n",i,value)
	}
	

	///CRUD:Create,Read,Update,Delete
	

	r := gin.Default()

	v1:=r.Group("/v1")
	{
		items:=v1.Group("/items")
		{
			items.POST("",ginitem.CreateItem(db))
			items.GET("",getListItem(db))
			items.GET("/:id",ginitem.GetItem(db))
			items.PATCH("/:id",ginitem.UpdateItem(db))
			items.DELETE("/:id",ginitem.DeleteItem(db))
		}
	}

  	r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Hello World",
    })
  })
  r.Run(PORT) // listen and serve on 0.0.0.0:3000(for windows "localhost:3000")


	
}



func getListItem(db *gorm.DB) func(*gin.Context){
	return func(c *gin.Context){

		var paging  common.Paging
		
		if err:=c.ShouldBind(&paging);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		paging.Process()

		var result [] models.TodoItem


		db = db.Where("status <> ?","Delete")

		if	err:=db.Table(models.TodoItem{}.TableName()).
		Count(&paging.Total).Error;err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		if err:=db.Order("id desc").
		Offset((paging.Page-1)*paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error;err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		

		c.JSON(http.StatusOK,common.NewSuccessResponse(result,paging,nil))
	}
}