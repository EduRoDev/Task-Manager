package main

import (
	"log"

	"github.com/EduRoDev/TaskManager/config"
	"github.com/EduRoDev/TaskManager/routes"
	"github.com/gin-contrib/cors"
)

func main() {
	config.InitDB()
	r := routes.SetupRoutes()
	log.Println("Server running on port 8080...")

	r.Use(cors.Default())
	r.Run(":8080")

}