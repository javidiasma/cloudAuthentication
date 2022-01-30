package main

import (
	"cloudAuthentication/config"
	"cloudAuthentication/routes"
	"fmt"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRoutes()
	err := r.Run("0.0.0.0:8001")

	if err != nil {
		fmt.Println(err.Error())
	}
}
