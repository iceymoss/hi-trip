package router

import (
	"hi-trip/trip_web/user_web/api"

	"github.com/gin-gonic/gin"
)

func InitUser(router *gin.RouterGroup) {
	router = router.Group("user")
	{
		router.POST("login", api.PassWordLogin)
		router.GET("list", api.List)
		router.POST("register", api.Register)
	}
}
