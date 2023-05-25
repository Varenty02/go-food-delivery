package main

import (
	"g05-fooddelivery/component/appctx"
	"g05-fooddelivery/middleware"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
	"g05-fooddelivery/module/restaurant/transport/ginrestaurant"
	"g05-fooddelivery/module/restaurantlike/transport/ginrstlike"
	ginupload "g05-fooddelivery/module/upload/transport"
	"g05-fooddelivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func setupRoute(ctx appctx.AppContext, v1 *gin.RouterGroup) {

	restaurant := v1.Group("/restaurants", middleware.RequireAuth(ctx))
	restaurant.POST("/", middleware.Recover(ctx), ginrestaurant.CreateRestaurant(ctx))

	//get
	restaurant.GET("/:id", ginrestaurant.FindRestaurant(ctx))
	//getall
	restaurant.GET("", ginrestaurant.ListRestaurant(ctx))
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
		ctx.GetMainDBConnection().Where("id=?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"restaurant": data,
		})
	})

	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(ctx))
	restaurant.POST("/:id/like", ginrstlike.UserLikeRestaurant(ctx))
	restaurant.DELETE("/:id/like", ginrstlike.UserUnlikeRestaurant(ctx))
	restaurant.GET("/:id/like", ginrstlike.ListUser(ctx))
	v1.POST("/register", ginuser.Register(ctx))
	v1.POST("/upload", ginupload.UploadImage(ctx))
	v1.POST("/authenticate", ginuser.Login(ctx))
	v1.GET("/profile", middleware.RequireAuth(ctx), ginuser.Profile(ctx))
}
