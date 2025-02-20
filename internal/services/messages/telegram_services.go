package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EduRoDev/TaskManager/config"
	"github.com/EduRoDev/TaskManager/internal/models"
	"github.com/joho/godotenv"
)

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendTelegramNotificacion(message string) error {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error al cargar el archivo .env: %v", err)
		return err
	}

	telegramToken := os.Getenv("TelegramBotToken")
	chatIdToken := os.Getenv("TelegramChatId")

	if telegramToken == "" || chatIdToken == "" {
		return fmt.Errorf("TelegramBotToken o TelegramChatId no están definidos en el archivo .env")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramToken)
	payload := TelegramMessage{
		ChatID: chatIdToken,
		Text:   message,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error al convertir el payload a JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error al enviar la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error al enviar el mensaje por Telegram, Status: %d", resp.StatusCode)
	}

	return nil
}

func CheckDueTaskandSendTelegram() {
	var tasks []models.Task
	result := config.Db.Where("notified_telegram = ?", false).Find(&tasks)

	if result.Error != nil {
		log.Println("❌ Error al obtener las tareas a vencer:", result.Error)
		return
	}

	for _, task := range tasks {
		message := fmt.Sprintf("📌 Recordatorio: Tu tarea '%s' vence pronto! ⏳", task.Title)

		err := SendTelegramNotificacion(message)

		if err != nil {
			log.Println("❌ Error al enviar el Telegram:", err)
		} else {
			log.Println("✅ Telegram enviado con éxito para la tarea:", task.Title)
			task.NotifiedTelegram = true
            config.Db.Save(&task)
		}
	}
}

