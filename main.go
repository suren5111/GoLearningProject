package main

import (
	"StudentService/cache"
	"StudentService/config"
	"StudentService/controller"
	"StudentService/database"
	"StudentService/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.InitDB(cfg)
	cache.InitRedis(cfg.Redis)

	r := gin.Default()

	// 登录和注销接口（不需要鉴权）
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)

	// 受保护的路由组
	protected := r.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/students", controller.ListStudents)
		protected.GET("/students/:id", controller.GetStudent)
		protected.POST("/students", controller.CreateStudent)
		protected.PUT("/students/:id", controller.UpdateStudent)
		protected.DELETE("/students/:id", controller.DeleteStudent)
	}

	r.Run(":8080")
}
