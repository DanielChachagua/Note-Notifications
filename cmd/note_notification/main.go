package main

import (
	"log"
	"note_notifications/cmd/note_notification/commands"
	"note_notifications/internal/database"
	"note_notifications/internal/dependencies"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/joho/godotenv"
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

	err = beeep.Notify("Alerta desde Go", "¡Tu programa CLI ha ejecutado una acción importante!", "")
	if err != nil {
		log.Fatalf("Error al enviar la notificación: %v", err)
	}

	deps := dependencies.NewContainer(db)

	commands.Execute(deps)
	
}