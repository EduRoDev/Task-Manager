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
    App := gin.Default()
   

    // Configurar CORS
    App.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "")
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
    App.GET("/tasks", taskController.GetAllTasksController)
    App.GET("/tasks/user/:userID", taskController.GetTasksByUserIDController) 
    App.POST("/tasks", taskController.CreateTaskHandler)
    App.PUT("/tasks/:id", taskController.UpdateTaskHandler)
    App.DELETE("/tasks/:id", taskController.DeleteTaskController)

    // Rutas de usuario
    App.POST("/user/login", userController.Login)
    App.POST("/user/register", userController.Register)
    App.POST("/user/ForgotPass",userController.ForgotPasswordHandler)
    App.POST("/user/ResetPass",userController.ResetPasswordHandler)

    
    return App
}


