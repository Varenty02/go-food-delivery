package main

import (
	"fmt"
	"g05-fooddelivery/module/component/appctx"
	"g05-fooddelivery/module/restaurant/middleware"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
	"g05-fooddelivery/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	//dunfg biến environment để lay duong dan
	dsn := os.Getenv("MYSQL_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()
	appctx := appctx.NewAppContext(db)
	fmt.Println("Hiếu đẹp trai")
	//giống fmt nhưng mà để log
	log.Println(db, err)
	//if err == nil {
	//	log.Fatalln(err)
	//}
	//Lấy 1 cái server
	r := gin.Default()
	r.Use(middleware.Recover(appctx))
	//Đăng ký đường link /ping cho server
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//post
	//v1 := r.Group("/v1", middleware.Recover(appctx))
	v1 := r.Group("/v1")
	restaurant := v1.Group("/restaurants")
	restaurant.POST("/", middleware.Recover(appctx), ginrestaurant.CreateRestaurant(appctx))

	//get
	restaurant.GET("/:id", ginrestaurant.FindRestaurant(appctx))
	//getall
	restaurant.GET("", ginrestaurant.ListRestaurant(appctx))
	//update
	restaurant.PUT("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		var data restaurantmodel.RestaurantUpdate
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

	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appctx))
	r.Run()
}
