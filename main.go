package main

import (
	"github.com/gin-gonic/gin"
	"udemy.com/rest-api/db"
	"udemy.com/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8200") //localhost:8200
}




