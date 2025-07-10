package middleware

import (
	"StudentService/cache"
	"StudentService/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTAuthMiddleware 验证 JWT 和 Redis 中的 token
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少 Authorization 头"})
			c.Abort()
			return
		}

		// 验证 token 是否存在于 Redis
		if !cache.IsTokenValid(tokenStr) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token 已过期或无效"})
			c.Abort()
			return
		}

		// 解析 JWT 获取用户信息
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token 解析失败"})
			c.Abort()
			return
		}

		// 将用户信息存入上下文，供后续使用
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
