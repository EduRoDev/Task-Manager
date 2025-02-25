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

func (uc *UserController) ForgotPasswordHandler(c *gin.Context) {
    var request struct {
        Email string `json:"email"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if err := uc.UserServices.ForgotPassword(request.Email); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password reset token sent"})
}

func (uc *UserController) ResetPasswordHandler(c *gin.Context) {
    var request struct {
        Email       string `json:"email"`
        Token       string `json:"token"`
        NewPassword string `json:"new_password"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if err := uc.UserServices.ResetPassword(request.Email, request.Token, request.NewPassword); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
}

