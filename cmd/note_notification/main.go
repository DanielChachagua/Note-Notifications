package main

import (
	"log"
	"note_notifications/cmd/note_notification/commands"
	"note_notifications/cmd/note_notification/jobs"
	"note_notifications/internal/database"
	"note_notifications/internal/dependencies"
	"note_notifications/internal/schemas"
	"os"
	"runtime"

	// "github.com/gen2brain/beeep"
	"github.com/joho/godotenv"
	// "github.com/robfig/cron/v3"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	db, err := database.ConectDB(os.Getenv("URI_DB"))
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	defer database.CloseDB(db)

	noteChan := make(chan []schemas.NoteResponse)

	runtime.LockOSThread()

	deps := dependencies.NewContainer(db)

	go func() {
		for notes := range noteChan {
			jobs.ShowAllNotesInOneWindow(&notes) // ✅ SE EJECUTA EN EL MAIN THREAD
		}
	}()

	jobs.InitNotifications(deps, noteChan)

	commands.Execute(deps)

	select {}
}
