package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("app starting ..")

	r := gin.New()
	r.Run()
}