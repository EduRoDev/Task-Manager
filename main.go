package main

import (
	"log"

	"github.com/EduRoDev/TaskManager/config"
	"github.com/EduRoDev/TaskManager/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := routes.SetupRoutes()
	log.Println("Server running on port 8080...")

	templates()
	r.Use(cors.Default())
	r.Run(":8080")

}

func templates(){
	App := gin.Default()
	templatesPath := "templates/**/*.html"
	App.LoadHTMLGlob(templatesPath)
	App.Static("/static", "./static")

	App.GET("/", func(c *gin.Context) {
		c.HTML(200,"Login.html",nil)
	})

	App.GET("/register", func(c *gin.Context) {
        c.HTML(200,"register.html",nil)
    })

    App.GET("/forgot", func(c *gin.Context) {
        c.HTML(200,"user/ForgotPassword.html",nil)
    })

	App.Run(":8080")
}

