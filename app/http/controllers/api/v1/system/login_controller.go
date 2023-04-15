package system

import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/system/sys_user"
	"gohub/app/requests"
	"gohub/app/requests/dto"
	"gohub/pkg/captcha"
	"gohub/pkg/jwt"
	"gohub/pkg/logger"
	"gohub/pkg/response"
)

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) Captcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)
	// 返回给用户
	response.Data(c, gin.H{
		"captcha_id":         id,
		"captcha_base64_img": b64s,
	})

}

func (lc *LoginController) Login(c *gin.Context) {
	request := dto.LoginByPasswordDto{}
	if ok := requests.Validate(c, &request, dto.ValidLoginByPassword); !ok {
		return
	}
	sysUser := sys_user.GetUser(request.UserName)
	if sysUser.ID == 0 || !sysUser.ComparePassword(request.Password) {
		response.Abort404(c, "账号不存在或密码错误")
	}

	// 登录成功
	token := jwt.NewJWT().IssueToken(sysUser.GetStringID(), sysUser.UserName)

	response.JSON(c, gin.H{
		"token": token,
	})
}
