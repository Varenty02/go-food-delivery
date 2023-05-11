package ginrestaurant

import (
	restaurantbiz "g05-fooddelivery/module/restaurant/biz"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
	restaurantstorage "g05-fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		//đổ dữlieeueeuj vào data
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSQLStore(db)

		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		biz.CreateRestaurant(c.Request.Context(), &data)
		c.JSON(http.StatusOK, gin.H{
			"restaurant": data,
		})
	}
}
