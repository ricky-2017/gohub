package dto

import (
	"gohub/app/requests"
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPasswordDto struct {
	UserName      string `json:"account,omitempty" valid:"user_name"`
	Password      string `json:"password,omitempty" valid:"password"`
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
}

func ValidLoginByPassword(data interface{}, c *gin.Context) map[string][]string {
	rule := govalidator.MapData{
		"user_name":      []string{"required", "min:3"},
		"password":       []string{"required", "min:6"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	message := govalidator.MapData{
		"user_name": []string{
			"required:账户为必填项",
			"min:账户ID长度需要超过3位",
		},
		"captcha_id": []string{
			"required:图形验证码不能为空",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}

	errs := requests.MakeValidate(data, rule, message)
	// 图片验证码
	_data := data.(*LoginByPasswordDto)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
