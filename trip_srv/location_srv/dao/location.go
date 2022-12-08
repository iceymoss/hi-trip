package dao

import (
	"hi-trip/global"
	"hi-trip/trip_srv/location_srv/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
//CreateLocation 新建搜索地点接口
func CreateLocation(req LocationInfo) (*LocationInfo, error){
	var location model.Location
	location.KayWord = req.KeyWord
	location.Latitude = req.Latitude
	location.Longitude = req.Longitude
	location.Boundary = req.Boundary
	location.UserID = req.UserID

	tx := global.DB.Save(&location)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.Aborted,  "新建位置失败")
	}
	return &LocationInfo{ID: location.ID}, nil

}

//LocationList 查询搜索历史
func LocationList(req LocationInfo) ([]LocationInfo, error){
	var list []model.Location
	tx := global.DB.Where(model.Location{UserID: req.UserID}).Find(&list)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "获取位置列表失败")
	}
	var infoList []LocationInfo
	for _, value := range list {
		info := LocationInfo{
			ID: value.ID,
			UserID: value.UserID,
			KeyWord: value.KayWord,
			Latitude: value.Latitude,
			Longitude: value.Longitude,
			Boundary: value.Boundary,
		}
		infoList = append(infoList, info)
	}

	return infoList, nil
}

//CreateKeyWord 新建关键字搜索记录
func CreateKeyWord(req LocationInfo)(*LocationInfo, error){
	var KeyWord model.KeywordSrv

	KeyWord.Region = req.Region
	KeyWord.KayWord = req.KeyWord
	KeyWord.UserID = req.UserID
	KeyWord.Latitude = req.Latitude
	KeyWord.Longitude = req.Longitude

	tx := global.DB.Save(&KeyWord)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.Aborted,  "新建位置失败")
	}

	return &LocationInfo{ID: KeyWord.ID}, nil
}

//KeyWordList 查询搜索历史
func KeyWordList(req LocationInfo) ([]LocationInfo, error){
	var list []model.KeywordSrv
	tx := global.DB.Where(model.Location{UserID: req.UserID}).Find(&list)
	if tx.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "获取位置列表失败")
	}
	var infoList []LocationInfo
	for _, value := range list {
		info := LocationInfo{
			ID: value.ID,
			UserID: value.UserID,
			KeyWord: value.KayWord,
			Latitude: value.Latitude,
			Longitude: value.Longitude,
			Region: value.Region,
		}
		infoList = append(infoList, info)
	}
	return infoList, nil
}





