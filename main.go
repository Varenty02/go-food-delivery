package main

import (
	"context"
	"fmt"
	"g05-fooddelivery/component/appctx"
	"g05-fooddelivery/middleware"
	"g05-fooddelivery/pubsub/localpb"
	"g05-fooddelivery/subscriber"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Địa chỉ và cổng Redis server
		Password: "",               // Mật khẩu Redis server (nếu có)
		DB:       0,                // Chọn cơ sở dữ liệu Redis
	})

	// Kiểm tra kết nối đến Redis server
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	if err := rdb.Set(context.Background(), "key", "value", 0).Err(); err != nil {
		panic(err)
	}

	// Truy xuất dữ liệu từ Redis
	value, err := rdb.Get(context.Background(), "key").Result()
	if err == redis.Nil {
		fmt.Println("Không tìm thấy key")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("Value:", value)
	}

	// Xóa dữ liệu từ Redis
	err = rdb.Del(context.Background(), "key").Err()
	if err != nil {
		panic(err)
	}
	//dsn := "root:123456@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	//dunfg biến environment để lay duong dan
	dsn := os.Getenv("MYSQL_CONN")
	secretKey := os.Getenv("SYSTEM_SECRET")
	ps := localpb.NewPubSub()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()
	appctx := appctx.NewAppContext(db, secretKey, ps)
	//subscriber.Setup(appctx, context.Background())
	_ = subscriber.NewEngine(appctx).Start()
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
