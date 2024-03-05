package config

import (
	"log"

	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/routes"
)

func Init() {
	config, err := database.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables 🚀", err)
	}
	database.ConnectDB(&config)
	r := routes.SetupRouter()
	//running the server
	r.Run(":" + config.ServerPort)
	log.Fatal(routes.SetupRouter().Run(":" + config.ServerPort))
}
