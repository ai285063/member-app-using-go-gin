package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	// router := gin.Default()

	ConnectMysql()
	ConnectRedis()

	g1 := router.Group("/users").Use(middleware)
	g1.GET("/", GetUserList)
	g1.POST("/register", Register)
	g1.PUT("/:id", PutUser)
	g1.DELETE("/:id", DeleteUser)

	g2 := router.Group("/viewcount")

	g2.GET("/", GetViewCount)

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}

func middleware(c *gin.Context) {
	c.Next()
	AddViewCount()
}
