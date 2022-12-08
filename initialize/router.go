package initialize

import (
	"net/http"

	"hi-trip/trip_web/router"
	"hi-trip/trip_web/router/location"

	"github.com/gin-gonic/gin"
)


//Routers 初始化及路由分发
func Routers() *gin.Engine {
	Router := gin.Default()

	ApiGroup := Router.Group("hi-trip")

	//分发路由
	ApiGroup = ApiGroup.Group("v1")

	router.InitUser(ApiGroup)
	router.InitSms(ApiGroup)
	location.InitSearch(ApiGroup)
	location.InitTans(ApiGroup)

	//健康检查
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "health",
		})
	})

	return Router
}
