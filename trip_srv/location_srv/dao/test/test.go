package main

import (
	"fmt"
	"hi-trip/initialize"
	"hi-trip/trip_srv/location_srv/dao"
)

func init() {
	initialize.InitializeDB()
}

func TestCreateLocation(){
	res, err := dao.CreateLocation(dao.LocationInfo{
		KeyWord: "北京",
		Longitude: 123.343213,
		Latitude: 120.342434,
		Boundary: 1000,
		UserID: 2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ID)
}

func TestList(){
	res, err := dao.LocationList(dao.LocationInfo{
		UserID: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestSearchKeyWord(){
	res, err := dao.CreateKeyWord(dao.LocationInfo{
		KeyWord: "火锅",
		Longitude: 123.343213,
		Latitude: 120.342434,
		Region: "无锡",
		UserID: 2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ID)
}

func TestKeyWordList(){
	res, err := dao.KeyWordList(dao.LocationInfo{
		UserID: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range res {
		fmt.Println(v)
	}

}

func main(){
	//for i := 0; i < 10; i++ {
	//	TestCreateLocation()
	//}

	//TestList()

	TestSearchKeyWord()
	TestKeyWordList()

}