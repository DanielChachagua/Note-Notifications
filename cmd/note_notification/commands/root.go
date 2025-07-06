package commands

import (
	"fmt"
	"note_notifications/cmd/note_notification/commands/calendar"
	"note_notifications/cmd/note_notification/commands/note"
	"note_notifications/internal/dependencies"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ntn",
	Short: "Una CLI para gestionar tus notas y notificaciones",
	Long:  "Una CLI para gestionar tus notas y notificaciones, ya sea creando, actualizando, listando o eliminando notas",
}

func Execute(deps *dependencies.Container) {
	// Añadir el comando 'note' y todos sus subcomandos al rootCmd
	rootCmd.AddCommand(note.NewNoteCmd(deps.Services.Note))
	rootCmd.AddCommand(calendar.NewCalendarCmd())

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
