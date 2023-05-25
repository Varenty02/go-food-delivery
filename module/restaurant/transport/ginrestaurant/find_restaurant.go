package ginrestaurant

import (
	"g05-fooddelivery/component/appctx"
	restaurantbiz "g05-fooddelivery/module/restaurant/biz"
	restaurantstorage "g05-fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	db := ctx.GetMainDBConnection()
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewFindRestaurantBiz(store)
		data, err := biz.FindDataWithCondition(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
