package requests

import (
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {
	rule := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}
	message := govalidator.MapData{
		"phone": []string{
			"required:手机号必填",
			"digits:手机号长度必须为11位数字",
		},
		"verify_code": []string{
			"required:验证码必填",
			"digits:验证码长度为6位",
		},
	}

	errs := validate(data, rule, message)

	// 手机验证码
	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}
