package router

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"pku-class/market/jwt"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, AuthorizeType")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func JWY() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/user/login/" {
			if c.GetHeader("AuthorizeType") == "password" {
				c.Next()
				return
			}
		}
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Illegal Token"})
				c.Abort()
				return
			}
			if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token Time Out"})
				c.Abort()
				return
			}
			c.Request.Header.Add("username", claims.Username)
			c.Request.Header.Add("id", strconv.Itoa(int(claims.Id)))
			c.Request.Header.Add("role", strconv.Itoa(claims.Role))
			log.Println("Token解析:")
			log.Println("用户名:" + claims.Username)
			log.Println("用户ID:" + strconv.Itoa(int(claims.Id)))
			log.Println("用户角色:" + strconv.Itoa(claims.Role))
		}
		c.Next()
		return
	}
}
