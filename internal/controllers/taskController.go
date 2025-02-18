package controllers

import (
	"net/http"
	"strconv"

	"github.com/EduRoDev/TaskManager/internal/models"
	"github.com/EduRoDev/TaskManager/internal/services"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskServices *services.TaskService
}

func (tc *TaskController) GetAllTasksController(c *gin.Context) {
    tasks, err := tc.TaskServices.GetAllTasks() 
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error fetching tasks"})
        return
    }
    c.JSON(http.StatusOK, tasks) 
}

func (tc *TaskController) CreateTaskHandler(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := tc.TaskServices.CreatTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating task"})
		return
	}

	go services.CheckDueTaskAndSendSMS()
	c.JSON(http.StatusCreated, task)
}


func (tc *TaskController) UpdateTaskHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := tc.TaskServices.UpdateTask(uint(id), &updatedTask); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	go services.CheckDueTaskAndSendSMS()
	c.JSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTaskController(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := tc.TaskServices.DeleteTask(uint(id)); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}