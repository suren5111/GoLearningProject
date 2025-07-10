package main

import (
	"StudentService/cache"
	"StudentService/config"
	"StudentService/controller"
	"StudentService/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.InitDB(cfg)
	cache.InitRedis(&cfg.Redis)

	r := gin.Default()
	r.GET("/students", controller.ListStudents)
	r.POST("/students", controller.CreateStudent)
	r.GET("/students/:id", controller.GetStudent)
	r.PUT("/students/:id", controller.UpdateStudent)
	r.DELETE("/students/:id", controller.DeleteStudent)

	r.Run(":8080")
}
