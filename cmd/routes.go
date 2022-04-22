package main

import (
	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "service is up",
		})
	})

	r.POST("/wallet/create/:userId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.GET("/wallet/getbalance/:userId", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.POST("/wallet/deposit/:userId/:amount", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.POST("/wallet/withdraw/:userId/:amount", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})		
	})

	return r
}