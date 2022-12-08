package dao

import (
	"crypto/sha512"
	"fmt"
	"strings"

	"hi-trip/global"
	"hi-trip/trip_srv/user_srv/model/user"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GetUserList 用户列表
func GetUserList() (*[]user.User, error) {
	var users []user.User

	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

//CreatUser 新建用户
func CreatUser(req *UserInfo) (int32, error) {
	var userinfo user.User
	if result := global.DB.Where(&user.User{Mobile: req.Mobile}).First(&userinfo); result.RowsAffected == 1 {
		return -1, status.Errorf(codes.AlreadyExists, "该手机号已注册")
	}

	//密码加密
	newPw := HashPassWord(req.PassWord)

	userInfo := user.User{
		NickName: req.NickName,
		Mobile:   req.Mobile,
		PassWord: newPw,
		Birthday: req.Birthday,
		Role:     1,
		Gender:   req.Gender,
		PerSig:   req.PerSig,
		Address:  req.Address,
		Email:    req.Email,
	}

	global.DB.Save(&userInfo)
	return userInfo.ID, nil
}

func MobileToUser(req *UserInfo) (*UserInfo, error) {
	var userInfo user.User
	if result := global.DB.Where(&user.User{Mobile: req.Mobile}).First(&userInfo); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "用户不存在")
	}

	rsp := UserInfo{
		ID:       userInfo.ID,
		NickName: userInfo.NickName,
		Mobile:   userInfo.Mobile,
		PassWord: userInfo.PassWord,
		Birthday: userInfo.Birthday,
		Gender:   userInfo.Gender,
		PerSig:   userInfo.PerSig,
		Address:  userInfo.Address,
	}

	return &rsp, nil
}

//UpdateUser 更新用户信息
func UpdateUser(req UserInfo) (*empty.Empty, error) {
	var user user.User
	result := global.DB.First(&user, req.ID)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "该用户不存在")
	}

	if req.Address != "" {
		user.Address = req.Address
	}

	if req.PerSig != "" {
		user.PerSig = req.PerSig
	}

	if req.Gender != "" {
		user.Gender = req.Gender
	}

	if req.Birthday != nil {
		user.Birthday = req.Birthday
	}

	if req.NickName != "" {
		user.NickName = req.NickName
	}

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	return &empty.Empty{}, nil
}

//DeleteUser 注销用户
func DeleteUser(req UserInfo) (*empty.Empty, error) {
	if result := global.DB.Delete(&user.User{}, req.ID); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	return &empty.Empty{}, nil

}

//HashPassWord 密码加密存储
func HashPassWord(pw string) string {
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodsPw := password.Encode(pw, options)

	NewPw := fmt.Sprintf("$hash521-blog$%s$%s", salt, encodsPw)
	return NewPw
}

//CheckPassWord 密码验证
func CheckPassWord(hashPw, Pw string) bool {
	PasswordInfo := strings.Split(hashPw, "$")
	fmt.Println(PasswordInfo)
	options := &password.Options{16, 100, 32, sha512.New}
	check := password.Verify(Pw, PasswordInfo[2], PasswordInfo[3], options)
	return check

}
