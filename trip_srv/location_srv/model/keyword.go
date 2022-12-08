package model


type KeywordSrv struct {
	BaseModel
	UserID int32
	Region string   `gorm:"type:varchar(50)"`
	KayWord string  `gorm:"type:varchar(100)"`
	Latitude float32 `gorm:"type:decimal(10,6);column:latitude"`
	Longitude float32  `gorm:"type:decimal(10,6);column:longitude"`
}

