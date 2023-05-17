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

func ListRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		pagingData.Fulfill()
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter.Status = []int{1}
		var result []restaurantmodel.Restaurant
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListDataWithCondition(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}
		//for i := range result {
		//	result[i].GenUID(common.DbtypeRestaurant)
		//}
		//for i := range result {
		//	result[i].Mask(false)
		//}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
