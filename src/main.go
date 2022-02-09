package main

import (
	"fmt"

	"github.com/dataset-license/portal-backend/src/config"
	"github.com/dataset-license/portal-backend/src/database"

	"github.com/dataset-license/portal-backend/src/routes"
)

func main() {
	if err := config.Load("src/config/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		return
	}

	_, err := database.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		return
	}

	router := routes.InitRouter()
	router.Run(config.Get().Addr)
}
