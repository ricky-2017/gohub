package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1 := r.Group("/v1")

	v1.Use(middlewares.LimitIP("200-H"))
	{
		// auth相关
		authGroup := v1.Group("/auth")

		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			// 判断手机是否注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitIP("60-H"), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitIP("60-H"), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			vfc := new(auth.VerifyCodeController)
			// 获取图形验证码
			authGroup.POST("/verfiy-codes/captcha", middlewares.GuestJWT(), vfc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.GuestJWT(), middlewares.LimitPerRoute("50-H"), vfc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.GuestJWT(), vfc.SendUsingEmail)

			// 用户登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByAccount)
			authGroup.POST("/login/refresh-token", middlewares.AuthJwt(), lgc.RefreshToken)

			ac := new(auth.AccountController)
			authGroup.POST("/reset-password/using-phone", middlewares.GuestJWT(), ac.ResetPasswordByPhone)
			authGroup.POST("/reset-password/using-email", middlewares.GuestJWT(), ac.ResetPasswordByEmail)
		}

	}
}
