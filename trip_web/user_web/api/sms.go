package api

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"hi-trip/trip_web/user_web/forms"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"go.uber.org/zap"
)

//GenerateSmsCode 生成width长度的验证码
func GenerateSmsCode(width int) string {
	number := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(number)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", number[rand.Intn(r)])
	}
	return sb.String()
}

//发送验证码
func SendSms(ctx *gin.Context) {
	//表单
	var forms forms.SendSmsForm
	if err := ctx.ShouldBind(&forms); err != nil {
		HandleValidatorErr(ctx, err)
		zap.S().Info("获取验证码表单失败", err)
		return
	}

	smsCode := GenerateSmsCode(5)

	credential := common.NewCredential(
		"SecretId",
		"SecretKey",
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 5
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	//实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	request := sms.NewSendSmsRequest()

	//短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666
	request.SmsSdkAppId = common.StringPtr("1400787878")

	//短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名
	request.SignName = common.StringPtr("腾讯云")

	//模板 ID: 必须填写已审核通过的模板 ID
	request.TemplateId = common.StringPtr("449739")

	//模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空
	request.TemplateParamSet = common.StringPtrs([]string{smsCode})

	//示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号
	request.PhoneNumberSet = common.StringPtrs([]string{"+8613711112222"})

	/* 用户的 session 内容（无需要可忽略）: 可以携带用户侧 ID 等上下文信息，server 会原样返回 */
	request.SessionContext = common.StringPtr("blog-go欢迎您注册")

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendSms(request)
	//TODO 处理异常

	_, _ = err.(*errors.TencentCloudSDKError)
	//TODO错误处理

	b, _ := json.Marshal(response.Response)
	// 打印返回的json字符串
	fmt.Printf("%s", b)

	//连接redis服务器
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
	})

	//写入缓存
	rdb.Set(context.Background(), forms.Mobile, smsCode, time.Second*time.Duration(600))

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
