package ginrestaurant

import (
	"fmt"
	"g05-fooddelivery/module/common"
	"g05-fooddelivery/module/component/appctx"
	restaurantbiz "g05-fooddelivery/module/restaurant/biz"
	restaurantmodel "g05-fooddelivery/module/restaurant/model"
	restaurantstorage "g05-fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered", r)
				}
			}()
			arr := []int{}
			log.Println(arr[0])
		}()
		arr := []int{}
		log.Println(arr[0])
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := restaurantstorage.NewSQLStore(db)

		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		//data.Mask(false)
		//c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
