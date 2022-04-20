package dto

import (
	"gohub/app/requests"
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetPasswordDto struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

// 验证器
func ValidResetPasswordByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required"},
		"password":    []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号不能为空",
			"digits:手机号长度必须为11位",
		},
		"verify_code": []string{
			"required:验证码不能为空",
		},
		"password": []string{
			"required:密码不能为空",
			"min:密码长度至少为6位",
		},
	}

	errs := requests.MakeValidate(data, rules, messages)

	// 验证码校验
	_data := data.(*ResetPasswordDto)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

// 通过邮箱验证码重置账户密码
type ResetPasswordByEmailDto struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

// 验证器
func ValidResetPasswordByEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":       []string{"required"},
		"verify_code": []string{"required"},
		"password":    []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:邮箱不能为空",
		},
		"verify_code": []string{
			"required:验证码不能为空",
		},
		"password": []string{
			"required:密码不能为空",
			"min:密码长度至少为6位",
		},
	}

	errs := requests.MakeValidate(data, rules, messages)

	// 验证码校验
	_data := data.(*ResetPasswordByEmailDto)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}
