package routes

import (
	"net/http"

	"github.com/EduRoDev/TaskManager/internal/controllers"
	"github.com/EduRoDev/TaskManager/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Middleware para permitir CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500") // Permitir solo este origen
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Manejar preflight request (OPTIONS)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	taskServices := &services.TaskService{}
	taskController := &controllers.TaskController{TaskServices: taskServices}

	r.GET("/tasks", taskController.GetAllTasksController)
	r.POST("/tasks", taskController.CreateTaskHandler)
	r.PUT("/tasks/:id", taskController.UpdateTaskHandler)
	r.DELETE("/tasks/:id", taskController.DeleteTaskController)

	return r
}
