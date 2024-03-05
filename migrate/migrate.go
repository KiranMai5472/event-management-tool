package main

import (
	"fmt"
	"log"

	"github.com/KiranMai5472/event-management-tool/database"
	"github.com/KiranMai5472/event-management-tool/models"
)

func init() {
	config, err := database.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	database.ConnectDB(&config)
}

func main() {
	database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	database.DB.AutoMigrate(&models.User{})
	fmt.Println("👍 Migration complete")
}
