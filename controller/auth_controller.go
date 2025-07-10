package controller

import (
	"StudentService/cache"
	"StudentService/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 登录接口
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	// 示例：验证用户名密码（可替换为数据库查询）
	if req.Username != "admin" || req.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 生成 JWT token
	token, err := utils.GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 token 失败"})
		return
	}

	// 将 token 存入 Redis，设置 10 分钟过期
	if err := cache.SetToken(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis 存储失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}

// 注销接口
func Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 Authorization 头"})
		return
	}

	// 从 Redis 删除 token
	if err := cache.DeleteToken(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注销失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注销成功"})
}
