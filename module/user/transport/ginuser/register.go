package ginuser

import (
	"g05-fooddelivery/common"
	"g05-fooddelivery/component/appctx"
	"g05-fooddelivery/component/hasher"
	userbiz "g05-fooddelivery/module/user/biz"
	usermodel "g05-fooddelivery/module/user/model"
	userstorage "g05-fooddelivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)
		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
