package ginuser

import (
	"g05-fooddelivery/common"
	"g05-fooddelivery/component/appctx"
	"g05-fooddelivery/component/hasher"
	"g05-fooddelivery/component/tokenprovider/jwt"
	userbiz "g05-fooddelivery/module/user/biz"
	usermodel "g05-fooddelivery/module/user/model"
	userstorage "g05-fooddelivery/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin
		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))

	}
}
