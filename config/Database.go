package config

import (
	"log"

	"github.com/EduRoDev/TaskManager/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	var err error
	dsn := "user=postgres password=123456 dbname=taskmanager port=5432 sslmode=disable TimeZone=America/Bogota"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	if err := Db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Error al migrar el modelo Task: %v", err)
	}

	log.Println("Conexión exitosa con la base de datos")
}