package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1 := r.Group("/v1").Use(middlewares.LimitIP("200-H"))
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
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			vfc := new(auth.VerifyCodeController)
			// 获取图形验证码
			authGroup.POST("/verfiy-codes/captcha", vfc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vfc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vfc.SendUsingEmail)

			// 用户登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			authGroup.POST("/login/using-password", lgc.LoginByAccount)
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)

			ac := new(auth.AccountController)
			authGroup.POST("/reset-password/using-phone", ac.ResetPasswordByPhone)
			authGroup.POST("/reset-password/using-email", ac.ResetPasswordByEmail)
		}

	}
}
