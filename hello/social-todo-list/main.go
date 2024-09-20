package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  

type TodoItem struct{
	Id int `json:"id" gorm:"column:id;"`
	Title string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status string `json:"status" gorm:"column:status;"`
	CreatedAt	*time.Time	`json:"created_at" gorm:"column:created_at;"`
	UpdatedAt	*time.Time	`json:"updated_at" gorm:"column:updated_at;"`
}

func (TodoItem)	TableName()	string{return "todo_items"}

type TodoItemCreation struct{
	Id	int	`json:"id" gorm:"column:id;"`
	Title	string	`json:"title" gorm:"column:title;"`
	Description	string	`json:"description"	gorm:"column:description;"`
}

type TodoItemUpdate struct{
	Title	*string	`json:"title" gorm:"column:title;"`
	Description	*string	`json:"description"	gorm:"column:description;"`
	Status *string `json:"status" gorm:"column:status;"`
}


func(TodoItemUpdate) TableName() string{return TodoItem{}.TableName()}



func (TodoItemCreation)	TableName()	string{return TodoItem{}.TableName()}


type Paging struct{
	Page int `json:"page" form:"page"`
	Limit	int	`json:"limit" form:"limit"`
	Total	int64 `json:"total" form:"-"`
}

func(p *Paging) Process(){
	if p.Page<=0{
		p.Page=1
	}

	if(p.Limit<=0||p.Limit>=100){
		p.Limit=10
	}
}

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

	

	now:=time.Now().UTC()

	item := TodoItem{
		Id:	1,
		Title:"This is item 1",
		Description:"This is item 1",
		Status:	"Doing",
		CreatedAt: &now,
		UpdatedAt: nil,
	}

	///Convert JSONData =byte[],err
	jsonData,err:= json.Marshal(item)


	if err!=nil{
		fmt.Println(err)
		return
	}

	///Convert to JSON
	fmt.Println(string(jsonData))




	var item2 TodoItem

	if err:=json.Unmarshal([]byte(jsonData),&item2); err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(item2)


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
			items.POST("",createItem(db))
			items.GET("",getListItem(db))
			items.GET("/:id",getItem(db))
			items.PATCH("/:id",updateItem(db))
			items.DELETE("/:id",deleteItem(db))
		}
	}

  	r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": item,
    })
  })
  r.Run(PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:3000")


	
}

func createItem(db	*gorm.DB) func(*gin.Context){
	return func(c *gin.Context){
		var data TodoItemCreation

		if err:=c.ShouldBind(&data);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		if err:=db.Create(&data).Error;err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		c.JSON(http.StatusOK,gin.H{
			"id":data.Id,
			"data":data,
		})
	}
}

func getItem(db *gorm.DB) func(*gin.Context){
	return func(c *gin.Context){
		var data  TodoItem

		id,err:=strconv.Atoi(c.Param("id"))

		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return 
		}


		if err:=db.Where("id = ?",id).First(&data).Error;err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		c.JSON(http.StatusOK,gin.H{
			"data":data,
		})
	}
}

func updateItem(db *gorm.DB) func(*gin.Context){
	return func(c *gin.Context){
		var data  TodoItemUpdate

		id,err:=strconv.Atoi(c.Param("id"))

		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return 
		}

		if err:=c.ShouldBind(&data);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}


		if err:=db.Where("id = ?",id).Updates(&data).Error;err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		c.JSON(http.StatusOK,gin.H{
			"data":true,
		})
	}
}


func deleteItem(db *gorm.DB) func(*gin.Context){
	return func(c *gin.Context){
		
		id,err:=strconv.Atoi(c.Param("id"))

		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return 
		}



		if err:=db.Table(TodoItem{}.TableName()).Where("id = ?",id).Updates(map[string]interface{}{
			"status":"Delete",
		}).Error;err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		c.JSON(http.StatusOK,gin.H{
			"data":true,
		})
	}
}

func getListItem(db *gorm.DB) func(*gin.Context){
	return func(c *gin.Context){

		var paging Paging
		
		if err:=c.ShouldBind(&paging);err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})

			return 
		}

		paging.Process()

		var result []TodoItem


		db = db.Where("status <> ?","Delete")

		if	err:=db.Table(TodoItem{}.TableName()).
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

		c.JSON(http.StatusOK,gin.H{
			"data":result,
			"paging":paging,
		})
	}
}