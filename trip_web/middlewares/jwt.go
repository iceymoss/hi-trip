package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"hi-trip/trip_web/user_web/models"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)


//JWTAuth 验证token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if err == TokenExpired {
					c.JSON(http.StatusUnauthorized, map[string]string{
						"msg": "授权已过期",
					})
					c.Abort()
					return
				}
			}

			c.JSON(http.StatusUnauthorized, "未登陆")
			c.Abort()
			return
		}

		//只要验证通过，需要证明你是谁,有没有权限
		c.Set("claims", claims)
		c.Set("userId", claims.ID)
		fmt.Println("token认证成功")
		c.Next()
	}
}

type JWT struct {
	singer []byte
}

func NewJWT() *JWT {
	return &JWT{
		singer: []byte("ice_moss"),
	}
}

//GenerateToken 生成token
func (j *JWT) GenerateToken(claims models.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	fmt.Println(token)
	newTokent, err := token.SignedString(j.singer)
	if err != nil {
		zap.S().Info("签名失败", err)
	}
	return newTokent, nil
}

//ParseToken 解析token
func (j *JWT) ParseToken(token string) (*models.Claims, error) {

	fmt.Println(token)

	tokenClaims, err := jwt.ParseWithClaims(token, &models.Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.singer, nil
	})

	if err != nil {
		fmt.Println("err:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if tokenClaims != nil {
		//断言
		if claims, ok := tokenClaims.Claims.(*models.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}

		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}

}
