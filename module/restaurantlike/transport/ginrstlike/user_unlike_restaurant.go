package ginrstlike

import (
	common2 "g05-fooddelivery/common"
	"g05-fooddelivery/component/appctx"
	reslikebiz "g05-fooddelivery/module/restaurantlike/biz"
	restaurantlikestorage "g05-fooddelivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserUnlikeRestaurant(appCxt appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requester := c.MustGet(common2.CurrentUser).(common2.Requester)
		resId, err := strconv.Atoi(id)
		if err != nil {
			panic(common2.ErrInvalidRequest(err))
		}
		store := restaurantlikestorage.NewSQLStore(appCxt.GetMainDBConnection())
		biz := reslikebiz.NewUserUnlikeRestaurantBiz(store)
		if err := biz.UnlikeRestaurant(c.Request.Context(), requester.GetUserId(), resId); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common2.SimpleSuccessResponse(true))
	}
}
