package model

import (
	"gorm.io/gorm"
	"time"
)

//BaseModel 公共字段
type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"` //column 定义别名
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type Location struct {
	BaseModel
	UserID int32
	KayWord string  `gorm:"type:varchar(100)"`
	Latitude float32 `gorm:"type:decimal(10,6);column:latitude"`
	Longitude float32  `gorm:"type:decimal(10,6);column:longitude"`
	Boundary  int32    `gorm:"column:boundary"`
}
