package main

import (
	"g05-fooddelivery/component/appctx"
	"g05-fooddelivery/middleware"
	"g05-fooddelivery/module/user/transport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupAdminRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin", middleware.RequireAuth(appContext))
	admin.GET("/profile", ginuser.Profile(appContext))
}
