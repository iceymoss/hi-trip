package router

import (
	"hi-trip/trip_web/user_web/api"

	"github.com/gin-gonic/gin"
)

func InitSms(router *gin.RouterGroup) {
	router = router.Group("base")
	{
		router.POST("send_sms", api.SendSms)
	}
}
