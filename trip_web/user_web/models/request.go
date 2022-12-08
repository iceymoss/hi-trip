package models

import "github.com/dgrijalva/jwt-go"

// Claims 鉴权结构
type Claims struct {
	ID       int64
	Username string
	AuthRole uint
	jwt.StandardClaims
}
