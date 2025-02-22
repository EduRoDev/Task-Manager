package services

import (
	"errors"

	"github.com/EduRoDev/TaskManager/config"
	"github.com/EduRoDev/TaskManager/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
    
}

func (us *UserServices) Register(email, password string) error {
    
    var user models.User
    if err := config.Db.Where("email = ?", email).First(&user).Error; err == nil {
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
    if err := config.Db.Create(&user).Error; err != nil {
        return err
    }

    return nil
}

func (us *UserServices) Login(email, password string) (*models.User, error) {
    
    var user models.User
    if err := config.Db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, errors.New("user not found")
    }

    
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("incorrect password")
    }

    return &user, nil
}

func (us *UserServices) EditPassword(email, password string) (*models.User,error){
    var user models.User
    if err := config.Db.Where("email = ?",email).Find(&user).Error; err != nil {
        return nil, errors.New("user not found")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user.Password = string(hashedPassword)
    if err := config.Db.Save(&user).Error; err != nil {
        return nil, err
    }
    
    return &user, nil
}