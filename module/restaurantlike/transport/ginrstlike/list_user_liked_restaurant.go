package ginrstlike

import (
	common2 "g05-fooddelivery/common"
	"g05-fooddelivery/component/appctx"
	reslikebiz "g05-fooddelivery/module/restaurantlike/biz"
	restaurantlikemodel "g05-fooddelivery/module/restaurantlike/model"
	restaurantlikestorage "g05-fooddelivery/module/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		resId, err := strconv.Atoi(id)
		if err != nil {
			panic(common2.ErrInvalidRequest(err))
		}
		filter := restaurantlikemodel.Filter{
			RestaurantId: resId,
		}
		var paging common2.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common2.ErrInvalidRequest(err))
		}
		paging.Fulfill()
		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := reslikebiz.NewListUserLikeRestaurantBiz(store)
		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common2.NewSuccessResponse(result, paging, filter))
	}
}
