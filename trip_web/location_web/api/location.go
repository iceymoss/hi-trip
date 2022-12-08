package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"hi-trip/trip_srv/location_srv/dao"
	"hi-trip/trip_web/location_web/forms"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

//HandleValidatorErr 表单验证错误处理返回
func HandleValidatorErr(c *gin.Context, err error) {

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": errs, //"errs.Translate(global.Trans)"
	})
}


func SearchLocation(ctx *gin.Context){
	var form forms.SearchLocationForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}

	tencentMapURL := ""

	//type为0位置搜索
	if form.SearchType == 0 {
		tencentMapURL = "https://apis.map.qq.com/ws/place/v1/search"
		var forminfo forms.SearchLocationForm
		if form.Boundary != 0 {
			forminfo.Boundary = form.Boundary
		}

		if form.Longitude != 0 && form.Latitude != 0 {
			forminfo.Longitude = form.Longitude
			forminfo.Latitude = form.Latitude
		}

		if form.Keyword != "" {
			forminfo.Keyword = form.Keyword
		}

		if form.PageSize != 0 {
			forminfo.PageSize = form.PageSize
		}

		if form.PageIndex != 0 {
			forminfo.PageIndex = form.PageIndex
		}

		tencentMapURL = fmt.Sprintf("%s?boundary=nearby(%f,%f,%d)&keyword=%s&page_size=%d&page_index=%d&key=MZQBZ-EKRKU-SK5VS-2DA2P-GIL6V-BUBEL",
			tencentMapURL, forminfo.Latitude, forminfo.Longitude, forminfo.Boundary, forminfo.Keyword, forminfo.PageSize, forminfo.PageIndex)

		fmt.Println(tencentMapURL)
	}else if form.SearchType == 1 {   //周边推荐
		tencentMapURL = "https://apis.map.qq.com/ws/place/v1/explore"
		var forminfo forms.SearchLocationForm
		if form.Boundary != 0 {
			forminfo.Boundary = form.Boundary
		}

		if form.Longitude != 0 && form.Latitude != 0 {
			forminfo.Longitude = form.Longitude
			forminfo.Latitude = form.Latitude
		}

		if form.PageSize != 0 {
			forminfo.PageSize = form.PageSize
		}

		if form.PageIndex != 0 {
			forminfo.PageIndex = form.PageIndex
		}

		tencentMapURL = fmt.Sprintf("%s?boundary=nearby(%f,%f,%d)&policy=1&page_size=%d&page_index=%d&key=MZQBZ-EKRKU-SK5VS-2DA2P-GIL6V-BUBEL",
			tencentMapURL, forminfo.Latitude, forminfo.Longitude, forminfo.Boundary, forminfo.PageSize, forminfo.PageIndex)

		fmt.Println(tencentMapURL)
	}else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "请选择搜索类型",
		})
		return
	}

	resp, err := http.Get(tencentMapURL)
	if err != nil {
		zap.S().Info("请求失败", err)
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"msg":"搜索失败",
		})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)



	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}

	userId, ok := ctx.Get("userId")
	if !ok {
		zap.S().Info("获取用户id失败")
	}

	userID := int32(userId.(int64))

	dbRes, err := dao.CreateLocation(dao.LocationInfo{
		UserID: userID,
		KeyWord: form.Keyword,
		Longitude: form.Longitude,
		Latitude: form.Latitude,
		Boundary: form.Boundary,
	})
	if err != nil {
		zap.S().Info("创建地点搜索记录失败", err)
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"msg":"搜索失败",
		})
		return
	}

	result := forms.SearchLocationResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("unmarshal失败", err)
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"msg":"搜索失败",
		})
		return
	}

	fmt.Println(result.DataInfo)

	ctx.JSON(http.StatusOK, gin.H{
		"id": dbRes.ID,
		"data": result,
	})
}

func List(ctx *gin.Context){
	userId, ok := ctx.Get("userId")
	if !ok {
		zap.S().Info("获取用户id失败")
	}

	userID := userId.(int64)

	Res, err := dao.LocationList(dao.LocationInfo{
		UserID: int32(userID),
	})
	if err != nil {
		zap.S().Info("查询搜索记录失败：", err)
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"msg":"获取搜索记录失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, Res)
}
