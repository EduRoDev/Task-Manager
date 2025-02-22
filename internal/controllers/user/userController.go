package controllers

import (
	"net/http"

	services "github.com/EduRoDev/TaskManager/internal/services/user"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserServices services.UserServices
}

type req struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func NewUserController(us *services.UserServices) *UserController{
	return &UserController{UserServices: *us}
}


func (usc *UserController) Register(c *gin.Context) {
	var request req
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := usc.UserServices.Register(request.Email,request.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	c.Status(http.StatusAccepted)
}

func (usc *UserController) Login(c *gin.Context){
	var request req 
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if user, err := usc.UserServices.Login(request.Email, request.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"login": "Accepted", "user": user})
	}
}

func (usc *UserController) EditPassword(c *gin.Context){
	var request req

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}


	user, err := usc.UserServices.EditPassword(request.Email,request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} 
	c.JSON(http.StatusOK,user)

}
