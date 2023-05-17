package ginrestaurant

import (
	"g05-fooddelivery/module/common"
	"g05-fooddelivery/module/component/appctx"
	restaurantbiz "g05-fooddelivery/module/restaurant/biz"
	restaurantstorage "g05-fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		id, err := strconv.Atoi(c.Param("id"))
		//uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)
		//if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
		//	panic(err)
		//} 
		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
