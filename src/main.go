package main

import (
	"fmt"

	"github.com/dataset-license/portal-backend/src/config"
	"github.com/dataset-license/portal-backend/src/database"

	"github.com/dataset-license/portal-backend/src/routes"
)

func main() {
	if err := config.Load(); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}
	fmt.Println("config is loaded")
	_, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}
	fmt.Println("DB is connected")
	router := routes.InitRouter()
	router.Run(config.Get().Addr)
}
