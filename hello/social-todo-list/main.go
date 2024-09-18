package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  

type TodoItem struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt	*time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at"`
}

func main(){


	// https://github.com/jackc/pgx
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(os.Getenv("APPNAME"))

	fmt.Println("Hello")

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
	

	r := gin.Default()
  	r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": item,
    })
  })
  r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")


	
}