package main

import (
	"golang/models"
	"golang/config"
	"golang/routes"
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Profile{})
	routes.SetRoutes(r)
	r.Run(":2000")
}