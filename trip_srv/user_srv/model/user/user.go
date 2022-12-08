package user

import (
	"time"

	"gorm.io/gorm"
)

//BaseModel 公共字段
type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"` //column 定义别名
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

//User 用户表单
type User struct {
	BaseModel
	NickName string     `gorm:"type: varchar(20) comment '昵称' "`
	Mobile   string     `gorm:"idx_mobile;unique type: varchar(11) comment '电话号码';not null"`
	PassWord string     `gorm:"type: varchar(100) comment '加密后的密码';not null"`
	Birthday *time.Time `gorm:"type:datetime"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示用户， 2表示管理员'"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'male表示男， famale表示女'"`
	PerSig   string     `gorm:"varchar(150) comment '个人签名'"`
	Address  string     `gorm:"varchar(20) comment '居住地址'"`
	Email    string     `gorm:"varchar(100) comment '邮件'"`
}
