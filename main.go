package main

import (
	"example.com/eventbooking/models/db"
	"example.com/eventbooking/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
