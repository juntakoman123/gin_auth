package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juntakoman123/gin_auth/controllers"
	"github.com/juntakoman123/gin_auth/middlewares"
	"github.com/juntakoman123/gin_auth/models"
)

func main() {
	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	router.Run(":8080")
}
