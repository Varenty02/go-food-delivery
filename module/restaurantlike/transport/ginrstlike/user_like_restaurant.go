package ginrstlike

import (
	common2 "g05-fooddelivery/common"
	"g05-fooddelivery/component/appctx"
	restaurantstorage "g05-fooddelivery/module/restaurant/storage"
	reslikebiz "g05-fooddelivery/module/restaurantlike/biz"
	restaurantlikemodel "g05-fooddelivery/module/restaurantlike/model"
	restaurantlikestorage "g05-fooddelivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		requester := c.MustGet(common2.CurrentUser).(common2.Requester)
		resId, err := strconv.Atoi(id)
		if err != nil {
			panic(common2.ErrInvalidRequest(err))
		}
		data := restaurantlikemodel.Like{
			RestaurantId: resId,
			UserId:       requester.GetUserId(),
		}
		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := reslikebiz.NewUserLikeRestaurantBiz(store, incStore)
		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common2.SimpleSuccessResponse(true))
	}
}
