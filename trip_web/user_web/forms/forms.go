package forms

type LoginForm struct {
	Mobile   string `from:"mobile" json:"mobile" binding:"required"` //电话号码有什么规律可寻，需要自定义validator
	Password string `from:"password" json:"password" binding:"required,min=3,max=20"`
}

type RegisterForm struct {
	Mobile   string `from:"mobile" json:"mobile" binding:"required"`
	NickName string `form:"nick_name" json:"nick_name" binding:"required"`
	Password string `from:"password" json:"password" binding:"required,min=3,max=20"`
	Code     string `from:"code" json:"code" binding:"required,min=5,max=5"`
}

type SendSmsForm struct {
	Mobile string `from:"mobile" json:"mobile" binding:"required"`       //电话号码有什么规律可寻，需要自定义validator
	Type   uint   `from:"type" json:"type" binding:"required,oneof=1 2"` //1表示注册，2表示登录 需要使用type区别，验证码的业务类别
}
