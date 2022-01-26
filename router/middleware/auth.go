package middleware

import (
	"log"
	"yoyo-mall/handler"
	"yoyo-mall/pkg/errno"
	"yoyo-mall/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		if err != nil {
			handler.SendUnauthorized(c, errno.ErrTokenInvalid, nil, err.Error())
			c.Abort()
			return
		}
		c.Set("id", ctx.ID)
		c.Set("role", ctx.Role)

		c.Next()
	}
}

// 管理员认证
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		if err != nil {
			handler.SendUnauthorized(c, errno.ErrTokenInvalid, nil, err.Error())
			c.Abort()
			return
		}
		if ctx.Role != 1 {
			handler.SendUnauthorized(c, errno.ErrTokenInvalid, nil, "admin is required")
			c.Abort()
			return
		}
		c.Set("id", ctx.ID)
		c.Set("role", ctx.Role)

		c.Next()
	}
}

func VisitorAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		if err != nil {
			log.Println("Token is invalid. Entry visitor mode.")
		} else {
			c.Set("id", ctx.ID)
			c.Set("role", ctx.Role)
		}

		c.Next()
	}
}
