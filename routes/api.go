package routes

import (
	controllers "gohub/app/http/controllers/api/v1"
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

			// 验证码相关
			vfc := new(auth.VerifyCodeController)
			authGroup.POST("/verfiy-codes/captcha", vfc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("50-H"), vfc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vfc.SendUsingEmail)

			// 用户登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByAccount)
			authGroup.POST("/login/refresh-token", middlewares.AuthJwt(), lgc.RefreshToken)

			ac := new(auth.AccountController)
			authGroup.POST("/reset-password/using-phone", middlewares.GuestJWT(), ac.ResetPasswordByPhone)
			authGroup.POST("/reset-password/using-email", middlewares.GuestJWT(), ac.ResetPasswordByEmail)
		}

		uc := new(controllers.UsersController)
		// 获取当前用户信息
		v1.GET("/user", middlewares.AuthJwt(), uc.CurrentUser)

		// 获取用户列表信息
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
		}

		// 分类相关
		//cgc := new(controllers.CategoriesController)
		//cgcGroup := v1.Group("/categories")
		//{
		//	cgcGroup.GET("", middlewares.AuthJwt(), cgc.Index)
		//	cgcGroup.POST("", middlewares.AuthJwt(), cgc.Store)
		//	cgcGroup.PUT("/:id", middlewares.AuthJwt(), cgc.Update)
		//}

		//tpc := new(controllers.TopicsController)
		//tpcGroup := v1.Group("/topics")
		//{
		//	tpcGroup.POST("", middlewares.AuthJwt(), tpc.Store)
		//	tpcGroup.PUT("/:id", middlewares.AuthJwt(), tpc.Update)
		//	tpcGroup.DELETE("/:id", middlewares.AuthJwt(), tpc.Delete)
		//}

		////////////////////// 客户端路由 ////////////////////////////

		// 文章控制器路由
		articleController := new(controllers.ArticlesController)
		articleGroup := v1.Group("/article")
		{
			articleGroup.GET("", middlewares.CrossOrigin(), articleController.Lists)
			articleGroup.GET("/:id", middlewares.CrossOrigin(), articleController.Show)
		}

		// 文章标签
		articleTagController := new(controllers.TagsController)
		articleTagGroup := v1.Group("/tag")
		{
			articleTagGroup.GET("", middlewares.CrossOrigin(), articleTagController.All)
		}

		// 文章分类
		categoriesController := new(controllers.CategoriesController)
		categoriesGroup := v1.Group("/category")
		{
			categoriesGroup.GET("", middlewares.CrossOrigin(), categoriesController.All)
		}

	}
}
