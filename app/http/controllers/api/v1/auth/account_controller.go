package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/app/requests/dto"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	v1.BaseAPIController
}

func (ac *AccountController) ResetPasswordByPhone(c *gin.Context) {
	// 参数校验
	request := dto.ResetPasswordDto{}
	if ok := requests.Validate(c, &request, dto.ValidResetPasswordByPhone); !ok {
		return
	}
	// 2. 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}

func (ac *AccountController) ResetPasswordByEmail(c *gin.Context) {
	// 参数校验
	request := dto.ResetPasswordByEmailDto{}
	if ok := requests.Validate(c, &request, dto.ValidResetPasswordByEmail); !ok {
		return
	}
	// 2. 更新密码
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}
