package middleware

import (
	"github.com/gin-gonic/gin"
	"mall/pkg/e"
	"mall/pkg/util"
	"time"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		//http的authorization字段通常用于传递身份验证信息，包含客户端的凭证，以便于服务器能够验证用户的身份

		token := c.GetHeader("Authorization")
		if token == "" {
			//登陆失败
			code = 404
		} else {
			claims, err := util.ParseToken(token) //根据当前的token进行解析用户信息
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				//说明当前的token已经过期了
				code = e.ErrorExpiredToken
			}

		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
