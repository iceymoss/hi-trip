package dao

import "time"

//UserInfo info
type UserInfo struct {
	ID       int32
	NickName string
	Mobile   string
	PassWord string
	Birthday *time.Time
	Role     int
	Gender   string
	PerSig   string
	Address  string
	Email    string
}

type UserList struct {
	List  []UserInfo
	Total int
}
