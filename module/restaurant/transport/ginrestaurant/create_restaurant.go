package ginrestaurant

import (
	"g05-fooddelivery/module/common"
	"g05-fooddelivery/module/component/appctx"
	restaurantbiz "g05-fooddelivery/module/restaurant/biz"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
	restaurantstorage "g05-fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSQLStore(db)

		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
