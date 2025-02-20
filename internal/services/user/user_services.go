package services

import (
	"errors"

	"github.com/EduRoDev/TaskManager/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices struct {
    DB *gorm.DB
}

func (us *UserServices) Register(email, password string) error {
    
    var user models.User
    if err := us.DB.Where("email = ?", email).First(&user).Error; err == nil {
        return errors.New("user already exists")
    }

    
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    
    user = models.User{
        Email:    email,
        Password: string(hashedPassword),
    }
    if err := us.DB.Create(&user).Error; err != nil {
        return err
    }

    return nil
}

func (us *UserServices) Login(email, password string) (*models.User, error) {
    
    var user models.User
    if err := us.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, errors.New("user not found")
    }

    
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("incorrect password")
    }

    return &user, nil
}