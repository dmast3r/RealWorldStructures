package main

import (
	"fmt"
	"os"
	"token-bucket-rate-limiter/internal/routes"
)

func main() {
	router := routes.GetRouter()

	if err := router.Run(":" + "8080"); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
