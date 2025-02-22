package models

import (
	"time"
)

type Task struct {
    ID              uint      `json:"id" gorm:"primaryKey"`
    Title           string    `json:"title" gorm:"not null"`
    Description     string    `json:"description"`
    IsDone          bool      `json:"is_done"`
    NotifiedSms     bool      `json:"notified_sms"`
    NotifiedTelegram bool     `json:"notified_telegram"`
    UserID          uint      `json:"user_id"`
    User            User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}
