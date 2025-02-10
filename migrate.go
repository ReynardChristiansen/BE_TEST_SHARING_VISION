package main

import (
	"BE/config"
	"BE/models"
	"log"
)

func main() {
	config.CreateDatabase()

	config.ConnectDB()

	config.DB.AutoMigrate(&models.Posts{})
	log.Println("Migration completed!")
}
