package main

import (
	"log"
	"note_notifications/cmd/note_notification/commands"
	"note_notifications/cmd/note_notification/functions"
	"note_notifications/internal/database"
	"note_notifications/internal/dependencies"

	// "note_notifications/internal/schemas"
	"os"
	// "runtime"

	// "github.com/gen2brain/beeep"
	"github.com/joho/godotenv"
	// "golang.org/x/oauth2/google"
	// "github.com/robfig/cron/v3"
	// "github.com/gen2brain/dlgs"
)

func main() {
	err := functions.InitNotifications()
	if err != nil {
		log.Fatal(err)
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
