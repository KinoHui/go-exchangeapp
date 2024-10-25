package main

import (
	"exchangeapp/backend/config"
	"exchangeapp/backend/router"

	"fmt"
)

func main() {
	fmt.Println("######TEST######")
	config.InitConfig()
	r := router.SetUpRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	r.Run(port) // listen and serve on 0.0.0.0:8080
}
