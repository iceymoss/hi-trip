package location

import (
	"hi-trip/trip_web/location_web/api"
	"hi-trip/trip_web/middlewares"

	"github.com/gin-gonic/gin"

)

func InitSearch(router *gin.RouterGroup){
	SearchRouter := router.Group("search")
	{
		SearchRouter.POST("/point",middlewares.JWTAuth(), api.SearchLocation)
		SearchRouter.GET("/location_list",middlewares.JWTAuth(), api.List)
		SearchRouter.POST("/key_word", middlewares.JWTAuth(), api.SearchKeyWord)
		SearchRouter.GET("/word_list", middlewares.JWTAuth(), api.KeyWordList)
	}
}

func InitTans(router *gin.RouterGroup){
	TansLocation := router.Group("tans")
	{
		TansLocation.POST("/lait_long", api.AddressTransLaiAndLon)
		TansLocation.POST("/address", api.LaiAndLongTransAddress)
	}
}
