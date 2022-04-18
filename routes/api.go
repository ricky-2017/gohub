package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// 注册路由
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"1": "2",
			})
		})

		// auth相关
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			vfc := new(auth.VerifyCodeController)
			// 获取图形验证码
			authGroup.POST("/verfiy-codes/captcha", vfc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vfc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vfc.SendUsingEmail)
		}

	}
}
