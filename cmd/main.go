package main

import (
	"fmt"
)

func main() {
	fmt.Println("app starting ..")

	r := setupRoutes()
	fmt.Println("service listening on port 8080")
	r.Run()
}