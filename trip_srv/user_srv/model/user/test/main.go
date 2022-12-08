package main

import (
	"fmt"
	"hi-trip/initialize"
	"hi-trip/trip_srv/user_srv/dao"
	"log"
)

func init() {
	initialize.InitializeDB()
}

func TestList() {
	res, err := dao.GetUserList()
	if err != nil {
		log.Fatal("获取失败", err)
	}

	for _, v := range *res {
		fmt.Println("用户信息", v)

		secss := dao.CheckPassWord(v.PassWord, "admin123")
		if secss {
			fmt.Println("密码验证成功")
		} else {
			fmt.Println("密码错误")
		}
	}

}
func TestCreate() {
	res, err := dao.CreatUser(&dao.UserInfo{
		NickName: "ice_moss",
		Address:  "无锡",
		Email:    "ice_moss@163.com",
		Mobile:   "1758343422",
		PassWord: "admin123",
	})
	if err != nil {
		log.Fatal("新建失败", err)
	}

	fmt.Println(res)
}

func TestUpdate() {
	_, err := dao.UpdateUser(dao.UserInfo{
		ID:     1,
		PerSig: "this is a test",
	})
	if err != nil {
		log.Fatal("更新失败", err)
	}
}

func TestDelete() {
	_, err := dao.DeleteUser(dao.UserInfo{
		ID: 1,
	})
	if err != nil {
		log.Fatal("注销失败", err)
	}
}

func main() {
	TestList()
}
