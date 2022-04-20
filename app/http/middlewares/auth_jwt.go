package middlewares

import (
	"fmt"
	"gohub/app/models/user"
	"gohub/pkg/config"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// 用户登录JWT鉴权中间件
func AuthJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(ctx)

		if err != nil {
			response.Unauthorized(ctx, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(ctx, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		ctx.Set("current_user_id", userModel.GetStringID())
		ctx.Set("current_user_name", userModel.Name)
		ctx.Set("current_user", userModel)

		ctx.Next()
	}
}
