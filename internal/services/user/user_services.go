package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/EduRoDev/TaskManager/config"
	"github.com/EduRoDev/TaskManager/internal/models"
	"github.com/EduRoDev/TaskManager/internal/services/messages"
	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
    
}

func generateResetToken() (string, error) {
    bytes := make([]byte, 16)
    if _, err := rand.Read(bytes); err != nil {
        return "",err
    }
    return hex.EncodeToString(bytes), nil
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

func (us *UserServices) ForgotPassword(email string ) error{
    var user models.User
    if err := config.Db.Where("email = ?", email).First(&user).Error; err != nil{
        return errors.New("user not found")
    }

    token, err := generateResetToken()
    if err != nil{
        return err
    }

    user.ResetToken = token
    user.ResetTokenExpiry = time.Now().Add(1 * time.Hour).Format(time.RFC3339)
    if err := config.Db.Save(&user).Error; err != nil {
        return err
    }

    message := fmt.Sprintf("Para restablecer tu contraseña, usa el siguiente token: %s", token)
    if err := messages.SendResetPasswordToTelegram(message); err != nil {
        return err
    }

    return nil
    
}

func (us *UserServices) ResetPassword(email, token, newPassword string) error {
    var user models.User
    if err := config.Db.Where("email = ? AND reset_token = ?", email, token).First(&user).Error; err != nil {
        return errors.New("invalid token or email")
    }

    expiryTime, err := time.Parse(time.RFC3339, user.ResetTokenExpiry)
    if err != nil {
        return errors.New("invalid token expiry format")
    }
    if time.Now().After(expiryTime) {
        return errors.New("token expired")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user.Password = string(hashedPassword)
    user.ResetToken = ""
    user.ResetTokenExpiry = ""
    if err := config.Db.Save(&user).Error; err != nil {
        return err
    }

    return nil
}
