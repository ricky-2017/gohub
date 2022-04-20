package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/app/requests/dto"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	// 参数校验
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}
	// 登录尝试
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// 账号/邮箱/手机 登录
func (lc *LoginController) LoginByAccount(c *gin.Context) {
	request := dto.LoginByPasswordDto{}
	if ok := requests.Validate(c, &request, dto.ValidLoginByPassword); !ok {
		return
	}
	user, err := auth.LoginByAccount(request.Account, request.CaptchaID, request.CaptchaAnswer, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

// token刷新
func (lc *LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)
	if err != nil {
		response.Error(c, err, "token刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
