package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}