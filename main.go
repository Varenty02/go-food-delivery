package main

import (
	"fmt"
	"g05-fooddelivery/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id     int    `json:"id" gorm:"column:id;"`
	Name   string `json:"name" gorm:"column:name;"`
	Addr   string `json:"addr" gorm:"column:addr;"`
	Status int    `json:"status" gorm:"column:status;"`
}

// dành cho update để biết được giá trị có sự thay đổi không(gorm không nhận 0 "" false
type RestaurantUpdate struct {
	Name   *string `json:"name" gorm:"column:name;"`
	Addr   *string `json:"addr" gorm:"column:addr;"`
	Status *int    `json:"status" gorm:"column:status;"`
}

func (Restaurant) TableName() string       { return "restaurants" }
func (RestaurantUpdate) TableName() string { return "restaurants" }
func main() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	//dunfg biến environment để lay duong dan
	dsn := os.Getenv("MYSQL_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("Hiếu đẹp trai")
	//giống fmt nhưng mà để log
	log.Println(db, err)
	//if err == nil {
	//	log.Fatalln(err)
	//}
	//Lấy 1 cái server
	r := gin.Default()
	//Đăng ký đường link /ping cho server
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//post
	v1 := r.Group("/v1")
	restaurant := v1.Group("/restaurants")
	restaurant.POST("/", ginrestaurant.CreateRestaurant(db))

	//get
	restaurant.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data Restaurant
		db.Where("id= ?", id).First(&data)
		c.JSON(http.StatusOK, gin.H{
			"restaurant": data,
		})
	})
	//getall
	restaurant.GET("", func(c *gin.Context) {
		var data []Restaurant
		//kỹ thuật phân trang
		type Panging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}
		var pagingData Panging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}
		if pagingData.Limit <= 0 {
			pagingData.Limit = 5
		}
		db.Order("id desc").Offset((pagingData.Page - 1) * pagingData.Limit).Limit(pagingData.Limit).Find(&data)
		c.JSON(http.StatusOK, gin.H{
			"restaurants": data,
		})
	})
	//update
	restaurant.PUT("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		var data RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"err": err.Error(),
			})
		}
		db.Where("id=?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"restaurant": data,
		})
	})

	restaurant.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		db.Table(Restaurant{}.TableName()).Where("id=?", id).Delete(nil)
		c.JSON(http.StatusOK, gin.H{
			"delete": "success",
		})
	})
	r.Run()
}
