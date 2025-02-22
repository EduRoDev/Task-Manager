package routes

import (
	"net/http"

	taskController "github.com/EduRoDev/TaskManager/internal/controllers/task"
	userController "github.com/EduRoDev/TaskManager/internal/controllers/user"
	task "github.com/EduRoDev/TaskManager/internal/services/task"
	user "github.com/EduRoDev/TaskManager/internal/services/user"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()
   

    // Configurar CORS
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    })

    // Configurar servicios y controladores
    taskServices := &task.TaskService{}
    taskController := &taskController.TaskController{TaskServices: taskServices}

    userServices := &user.UserServices{}
    userController := &userController.UserController{UserServices: *userServices}

    // Rutas de tareas
    r.GET("/tasks", taskController.GetAllTasksController)
    r.POST("/tasks", taskController.CreateTaskHandler)
    r.PUT("/tasks/:id", taskController.UpdateTaskHandler)
    r.DELETE("/tasks/:id", taskController.DeleteTaskController)

    // Rutas de usuario
    r.POST("/user/login", userController.Login)
    r.POST("/user/register", userController.Register)
    r.PUT("/user/edit/password", userController.EditPassword)

    

    return r
}
