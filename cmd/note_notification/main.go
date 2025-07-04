package main

import (
	"fmt"
	"log"
	"note_notifications/cmd/note_notification/commands"
	"note_notifications/cmd/note_notification/jobs"
	"note_notifications/internal/database"
	"note_notifications/internal/dependencies"
	// "note_notifications/internal/schemas"
	"os"
	// "runtime"
	"time"

	// "github.com/gen2brain/beeep"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
	// "github.com/robfig/cron/v3"
	// "github.com/gen2brain/dlgs"
)

func main() {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("No se pudo leer el archivo credentials.json: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Error al parsear config: %v", err)
	}
	client := jobs.GetClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("No se pudo crear el cliente de Calendar: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("No se pudieron obtener los eventos: %v", err)
	}

	fmt.Println("Próximos eventos:")
	if len(events.Items) == 0 {
		fmt.Println("No hay eventos próximos.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v - %s\n", date, item.Summary)
		}
	}
	// confirmed, err := dlgs.Question("Apagar sistema", "¿Querés apagar la computadora?", true)
	// dlgs.Warning("Apagando sistema...", "Apagando sistema2...")
	// if err != nil {
	// 	panic(err)
	// }
	// item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	// if err != nil {
	// 	panic(err)
	// }
	// passwd, _, err := dlgs.Password("Password", "Enter your API key:")
	// if err != nil {
	// 	panic(err)
	// }

	// val, bol , err := dlgs.Entry("fecha", "Enter your API key:", "02-01-2006")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val, bol)

	// if item == "Bug" {
	// 	fmt.Println("Apagando sistema...")
	// 	panic("Apagando sistema...")
	// 	// Aquí ejecutás tu comando de apagado
	// }

	// if passwd == "1234" {
	// 	fmt.Println("Apagando sistema...")
	// 	panic("Apagando sistema...")
	// 	// Aquí ejecutás tu comando de apagado
	// }

	// if confirmed {
	// 	fmt.Println("Apagando sistema...")
	// 	panic("Apagando sistema...")
	// 	// Aquí ejecutás tu comando de apagado
	// } else {
	// 	fmt.Println("Cancelado")
	// }

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	db, err := database.ConectDB(os.Getenv("URI_DB"))
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer database.CloseDB(db)

	// noteChan := make(chan []schemas.NoteResponse)

	// runtime.LockOSThread()

	deps := dependencies.NewContainer(db)

	// go func() {
	// 	for notes := range noteChan {
	// 		jobs.ShowAllNotesInOneWindow(&notes) // ✅ SE EJECUTA EN EL MAIN THREAD
	// 	}
	// }()

	// jobs.InitNotifications(deps, noteChan)

	commands.Execute(deps)

	select {}
}
