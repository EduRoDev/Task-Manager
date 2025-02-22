package services

import (
	"errors"
	"log"

	"github.com/EduRoDev/TaskManager/config"
	"github.com/EduRoDev/TaskManager/internal/models"
)


type TaskService struct{}



func (ts *TaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := config.Db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}


func (ts *TaskService) CreatTask(task *models.Task) error {
	
	result := config.Db.Create(task)
	if result.Error != nil {
		log.Printf("Error al crear la tarea: %s", result.Error)
		return result.Error
	}
	return nil
}


func (ts *TaskService) UpdateTask(id uint,Updatetask *models.Task) error{
	var task models.Task
	if err := config.Db.First(&task, id).Error; err != nil{
		return errors.New("task not found")
	}

	task.Title = Updatetask.Title
	task.Description = Updatetask.Description
	task.IsDone = Updatetask.IsDone
	task.UserID = Updatetask.UserID

	result := config.Db.Save(&task)
	if result.Error != nil {
		log.Printf("Error al actualizar la tarea: %s", result.Error)
		return result.Error
	}
	return nil
}

func (ts *TaskService) DeleteTask(id uint) error{
	result := config.Db.Delete(&models.Task{}, id)
	if result.Error != nil {
		log.Printf("Error al eliminar la tarea: %s",result.Error)
		return result.Error
	}
	return nil
}