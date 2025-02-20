package messages

import (
	"fmt"
	"log"
	"os"

	"github.com/EduRoDev/TaskManager/config"
	"github.com/EduRoDev/TaskManager/internal/models"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)



func SendSMS(to string, message string) error {
	err := godotenv.Load("../.env")
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: os.Getenv("AccountSID"),
        Password: os.Getenv("AuthToken"),
    })

    params := &api.CreateMessageParams{}
    params.SetTo(to)
    params.SetFrom(os.Getenv("PhoneNumber"))
    params.SetBody(message)

    resp, err := client.Api.CreateMessage(params)
    if (err != nil) {
        log.Printf("Error al enviar SMS: %v", err)
        return err
    }

    fmt.Printf("✅ SMS enviado con éxito a %s, SID: %s\n", to, *resp.Sid)
    return nil
}

func CheckDueTaskAndSendSMS() {
    var tasks []models.Task
    result := config.Db.Where("notified_sms = ?", false).Find(&tasks)

    if result.Error != nil {
        log.Println("❌ Error al obtener las tareas a vencer:", result.Error)
        return
    }

    for _, task := range tasks {
        message := fmt.Sprintf("📌 Recordatorio: Tu tarea '%s' vence pronto! ⏳", task.Title)

        err := SendSMS("+573016672931", message)
        
        if err != nil {
            log.Println("❌ Error al enviar el SMS:", err)
        } else {
            log.Println("✅ SMS enviado con éxito para la tarea:", task.Title)

            task.NotifiedSms = true
            config.Db.Save(&task)
        }
    }
}
