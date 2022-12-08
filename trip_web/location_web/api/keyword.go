package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"hi-trip/trip_srv/location_srv/dao"
	"hi-trip/trip_web/location_web/forms"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SearchKeyWord(ctx *gin.Context){
	var form forms.SearchKeyWordForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}

	tencentMapURL := "https://apis.map.qq.com/ws/place/v1/suggestion/"

	tencentMapURL = fmt.Sprintf("%s?region=%s&keyword=%s&key=MZQBZ-EKRKU-SK5VS-2DA2P-GIL6V-BUBEL",tencentMapURL, form.Region, form.Keyword)

	res, err := http.Get(tencentMapURL)
	if err != nil {
		zap.S().Info("关键字搜索请求失败", err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"msg":"请求失败",
		})
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		zap.S().Info("读取至缓冲区失败", err)
		return
	}

	if res.StatusCode != 200 {
		zap.S().Info("请求失败", err)
		return
	}

	result := forms.SearchLocationResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		zap.S().Info("json解析失败", err)
		return
	}

	userId, ok := ctx.Get("userId");
	if !ok {
		zap.S().Info("获取userID失败", userId)
		return
	}
	userID := int32(userId.(int64))
	dbres, err := dao.CreateKeyWord(dao.LocationInfo{
		UserID: userID,
		KeyWord: form.Keyword,
		Region: form.Region,
	})
	if err != nil {
		zap.S().Info("创建关键字搜索记录失败", err)
		return
	}

	fmt.Println(result)

	ctx.JSON(http.StatusOK, gin.H{
		"id": dbres.ID,
		"data": result,
	})
}

func KeyWordList(ctx *gin.Context){
	userId, ok := ctx.Get("userId");
	if !ok {
		zap.S().Info("获取userID失败", userId)
		return
	}
	userID := int32(userId.(int64))

	res, err :=dao.KeyWordList(dao.LocationInfo{
		UserID: userID,
	})
	if err != nil {
		zap.S().Info("获取搜索记录失败", err)
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"msg": "获取搜索记录失败",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
