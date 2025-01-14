package routers

import (
	"RemoteUnlockServer/src/routers/api"
	"RemoteUnlockServer/src/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//r.LoadHTMLGlob("bin/templates/*")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.ftl", gin.H{
	//		"title": "Main website",
	//	})
	//})
	r.POST("/UploadUnlockFile", api.UploadUnlockFile)
	//r.GET("/DownloadUnlockFile", api.DownloadUnlockFile)
	return r
}
