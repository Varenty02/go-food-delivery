package ginupload

import (
	"fmt"
	"g05-fooddelivery/common"
	"g05-fooddelivery/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}
		if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%d%s", time.Now().Nanosecond(), fileHeader.Filename)); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
			Id:        0,
			Url:       "http://localhost:8080/static/" + fileHeader.Filename,
			Width:     0,
			Height:    0,
			CloudName: "local",
			Extension: "png",
		}))
	}
}
