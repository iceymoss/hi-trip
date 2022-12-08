package api

import (
	"context"
	"fmt"
	middlewares2 "hi-trip/trip_web/middlewares"
	"log"
	"net/http"
	"time"

	"hi-trip/trip_srv/user_srv/dao"
	form "hi-trip/trip_web/user_web/forms"
	"hi-trip/trip_web/user_web/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

//HandleValidatorErr 表单验证错误处理返回
func HandleValidatorErr(c *gin.Context, err error) {

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": errs, //"errs.Translate(global.Trans)"
	})
}

func List(ctx *gin.Context) {
	res, err := dao.GetUserList()
	if err != nil {
		ctx.JSON(http.StatusFound, gin.H{
			"msg": "获取用户列表失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_list": res,
	})
}

//登录
func PassWordLogin(ctx *gin.Context) {
	//表单
	var forms form.LoginForm
	if err := ctx.ShouldBind(&forms); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}

	//获取数据库密码
	res, err := dao.MobileToUser(&dao.UserInfo{
		Mobile: forms.Mobile,
	})
	if err != nil {
		log.Fatal("手机号没有注册", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号没有注册",
		})
		return
	}

	//验证密码
	isCheck := dao.CheckPassWord(res.PassWord, forms.Password)
	if !isCheck {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"password": "登录失败",
		})

	} else {

		//生成token
		j := middlewares2.NewJWT()
		expirOut := time.Now().Unix()
		claims := models.Claims{
			int64(res.ID),
			res.NickName,
			uint(res.Role),
			jwt.StandardClaims{
				NotBefore: expirOut,
				ExpiresAt: expirOut + 60*60*24*30, //30天过期
				Issuer:    "ice_moss",
			},
		}

		token, err := j.GenerateToken(claims)
		if err != nil {
			fmt.Println(forms)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "登录失败",
			})
			return
		}

		//返回token及有效时间
		ctx.JSON(http.StatusOK, gin.H{
			"id":        res.ID,
			"user_name": res.NickName,
			"token":     token,
			"expir":     expirOut,
		})
	}
}

func Register(ctx *gin.Context) {
	//表单
	var form form.RegisterForm
	if err := ctx.ShouldBind(&form); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
	})

	mobile := "17585610985"
	rsp := rdb.Get(context.Background(), mobile)
	code, err := rsp.Result()
	if err == redis.Nil {
		zap.S().Info("验证码错误", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码错误",
		})
		return
	} else {
		if form.Code != code {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "验证码错误",
			})
			return
		}

		//默认为普通用户
		Role := 1
		userId, err := dao.CreatUser(&dao.UserInfo{
			NickName: form.NickName,
			Mobile:   form.Mobile,
			PassWord: form.Password,
			Role:     Role,
		})
		if err != nil {
			zap.S().Errorw("[Register] 创建 【用户失败】", err.Error())
			HandleValidatorErr(ctx, err)
			return
		}

		//直接登录
		//生成token
		j := middlewares2.NewJWT()
		expirOut := time.Now().Unix()
		claims := models.Claims{
			int64(userId),
			form.NickName,
			uint(Role),
			jwt.StandardClaims{
				NotBefore: expirOut,
				ExpiresAt: expirOut + 60*60*24*30, //30天过期
				Issuer:    "ice_moss",
			},
		}

		token, err := j.GenerateToken(claims)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "登录失败",
			})
			return
		}

		//返回token及有效时间
		ctx.JSON(http.StatusOK, gin.H{
			"id":        userId,
			"user_name": form.NickName,
			"token":     token,
			"expir":     expirOut,
		})
	}

}
