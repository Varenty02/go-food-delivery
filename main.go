package main

import (
	"fmt"
	"g05-fooddelivery/component/appctx"
	"g05-fooddelivery/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	//dunfg biến environment để lay duong dan
	dsn := os.Getenv("MYSQL_CONN")
	secretKey := os.Getenv("SYSTEM_SECRET")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()
	appctx := appctx.NewAppContext(db, secretKey)
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
	r.Static("/static", "./static")
	v1 := r.Group("/v1")
	setupRoute(appctx, v1)
	setupAdminRoute(appctx, v1)
	//post
	//v1 := r.Group("/v1", middleware.Recover(appctx))

	r.Run()
}
