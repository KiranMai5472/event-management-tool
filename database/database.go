package database

import (
	"fmt"
	"log"

	"github.com/KiranMai5472/event-management-tool/Constants"
	"github.com/KiranMai5472/event-management-tool/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	//dsn := os.Getenv("DATABASE_URL")
	fmt.Println(config.DBName)
	logger.LogDebug("Enter in ConnectDB()", Constants.LogFields)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogError(err, Constants.LogFields)
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database 🚀")
}
