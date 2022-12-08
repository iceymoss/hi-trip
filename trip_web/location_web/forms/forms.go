package forms


type SearchLocationForm struct {
	SearchType int `from:"search_type" json:"search_type" binding:"required"`
	Keyword string `from:"keyword" json:"keyword" binding:"required"`
	Boundary int32 `from:"boundary" json:"boundary" binding:"required"`
	GetSubpois int `from:"get_subpois" json:"get_subpois"`//是否返回子地点：默认为0不返回，1：返回
	Latitude float32 `from:"latitude" json:"latitude" binding:"required"`
	Longitude float32 `from:"longitude" json:"longitude" binding:"required"`
	Filter  string `from:"filter" json:"filter"`
	Orderby  string `from:"orderby" json:"orderby"`
	PageSize  int `from:"page_size" json:"page_size"`
	PageIndex int `from:"page_index" json:"page_index"`
}


type SearchKeyWordForm struct {
	Keyword string `from:"keyword" json:"keyword" binding:"required"`
	Region string `from:"region" json:"region" binding:"required"`
	Latitude float32 `from:"latitude" json:"latitude"`
	Longitude float32 `from:"longitude" json:"longitude"`
}

type AddressTransLaiAndLonForm struct {
	Address string `form:"address" json:"address" binding:"required"`
}


type Location struct {
	Latitude float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

type ADInfo struct {
	Adcode int32 `json:"adcode"`
	Province string `json:"province"`
	City string `json:"city"`
	District string `json:"district"`
}


type Data struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Address string `json:"address"`
	Tel string `json:"tel"`
	Category string `json:"category"`
	Type int32    `json:"type"`            //值说明：0:普通POI / 1:公交车站 / 2:地铁站 / 3:公交线路 / 4:行政区划
	Location  Location `json:"location"`
	Distance float32   `json:"_distance"`   //距离，单位： 米，在周边搜索、城市范围搜索传入定位点时返回
	ADInfo ADInfo  `json:"ad_info"`     //行政区
}


type SearchLocationResponse struct {
	Status int  `json:"status"`
	Message string `json:"message"`
	count int    `json:"count"`
	RequestId  string `json:"request_id"`
	DataInfo []Data   `json:"data"`
	SubPoise []Data `json:"sub_pois"`  //子区域
}

type AddressComponents struct {
	Province string `json:"province"`
	City string `json:"city"`
	District string `json:"district"`
	Street string `json:"street"`
	StreetNumber string `json:"street_number"`
}

type AdInformation struct {
	Adcode  string `json:"adcode"`
}

type Result struct {
	Title string `json:"title"`
	Location Location `json:"location"`
	AddressComponents AddressComponents `json:"address_components"`
	AdInformation  AdInformation `json:"ad_info"`
}

type AddressTransLaiAndLongResponse struct {
	Status int  `json:"status"`
	Message string `json:"message"`
	Result Result `json:"result"`
	Similarity int `json:"similarity"`
	Reliability int `json:"reliability"`
	Level  int `json:"level"`
}

type locationResult struct {
	Address string `json:"address"`
	AddressComponent AddressComponents `json:"address_component"`
}

type LaiAndLongTransAddressResponse struct {
	Status int  `json:"status"`
	Message string `json:"message"`
	Result locationResult `json:"result"`
}


