package jwt

import (
	"net/http"
	"strings"
	"time"
	"wx-gin-master/pkg/e"
	"wx-gin-master/pkg/util"

	"github.com/gin-gonic/gin"
)

// jwt 中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var msg string
		var data interface{}

		code = e.OK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.Unauthorized
		} else {
			if !strings.HasPrefix(token, "Bearer ") {
				code = e.Unauthorized
			} else {
				token = token[len("Bearer "):]
				claims, err := util.ParseToken(token)
				if err != nil {
					code = e.Unauthorized
					msg = "Invalid token"
				} else if time.Now().Unix() > claims.ExpiresAt {
					code = e.Unauthorized
					msg = "Invalid token"
				}
				c.Set("AuthData", claims)
			}
		}

		if code != e.OK {
			if msg == "" {
				msg = e.GetMsg(code)
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})

			c.Abort()
			return
		}
		//上面的代码是 before_request
		c.Next()
		//后面的代码是after request
	}
}

// Admin 管理员验证
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var msg string
		var data interface{}

		code = e.OK
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.Unauthorized
		} else {
			if !strings.HasPrefix(token, "Bearer ") {
				code = e.Unauthorized
			} else {
				token = token[len("Bearer "):]
				claims, err := util.ParseAdmin(token)
				if err != nil {
					code = e.Unauthorized
					msg = "Invalid token"
				} else if time.Now().Unix() > claims.ExpiresAt {
					code = e.Unauthorized
					msg = "Invalid token"
				}
				// 将解析后的有效载荷claims重新写入gin.Context引用对象中
				c.Set("AuthData", claims)
			}

		}

		if code != e.OK {
			if msg == "" {
				msg = e.GetMsg(code)
			}
			c.JSON(code, gin.H{
				"code": code,
				"msg":  msg,
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()

	}
}
