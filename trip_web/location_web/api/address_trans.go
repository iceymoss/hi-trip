package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"hi-trip/trip_web/location_web/forms"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

)

func AddressTransLaiAndLon(ctx *gin.Context){
	var addressForm forms.AddressTransLaiAndLonForm
	if err := ctx.ShouldBind(&addressForm); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}

	tencentMapURL := fmt.Sprintf("https://apis.map.qq.com/ws/geocoder/v1/?address=%s&key=MZQBZ-EKRKU-SK5VS-2DA2P-GIL6V-BUBEL", addressForm.Address)
	res, err := http.Get(tencentMapURL)
	if err != nil {
		zap.S().Info("请求地址转换失败", err)
		return
	}

	fmt.Println("返回结果：", res.Body)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		zap.S().Info("读取至缓冲区失败", err)
		return
	}

	result := forms.AddressTransLaiAndLongResponse{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		zap.S().Info("json解析失败", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func LaiAndLongTransAddress(ctx *gin.Context){
	var locationForm forms.Location
	if err := ctx.ShouldBind(&locationForm); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}

	tencentMapURL := fmt.Sprintf("https://apis.map.qq.com/ws/geocoder/v1/?location=%f,%f&key=MZQBZ-EKRKU-SK5VS-2DA2P-GIL6V-BUBEL&get_poi=1", locationForm.Latitude, locationForm.Longitude)
	res, err := http.Get(tencentMapURL)
	if err != nil {
		zap.S().Info("请求地址转换失败", err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		zap.S().Info("读取至缓冲区失败", err)
		return
	}

	result := forms.LaiAndLongTransAddressResponse{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		zap.S().Info("json解析失败", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}


