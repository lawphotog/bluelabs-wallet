package main

import (
	"bluelabs/wallet/internal/repository"
	"bluelabs/wallet/internal/wallet"
	"strconv"

	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	r := gin.Default()

	client := CreateLocalClient()
	repository := repository.New(client)

	repository.Setup() //one off to set up dynamodb

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "service is up",
		})
	})

	r.POST("/wallet/create/:userId", func(c *gin.Context) {
		userId := c.Params.ByName("userId")
		wallet := wallet.New(repository)
		err := wallet.Create(userId)

		//or error status code. depends on what client wants and consistency with other services
		if err != nil {
			SendError(c, err)
			return
		}

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
		userId := c.Params.ByName("userId") //or Post Body
		amount := c.Params.ByName("amount") //or Post Body
		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			SendError(c, err)
			return
		}

		wallet := wallet.New(repository)
		err = wallet.Deposit(userId, int64(amountInt))
		if err != nil {
			SendError(c, err)
			return
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.POST("/wallet/withdraw/:userId/:amount", func(c *gin.Context) {
		userId := c.Params.ByName("userId")
		amount := c.Params.ByName("amount")
		amountInt, err := strconv.Atoi(amount)
		if err != nil {
			SendError(c, err)			
		}

		wallet := wallet.New(repository)
		err = wallet.Withdraw(userId, int64(amountInt))
		if err != nil {
			SendError(c, err)
			return
		}
		c.JSON(200, gin.H{
			"message": "success",
		})		
	})

	return r
}

func SendError(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"message": err.Error(),
	})
}