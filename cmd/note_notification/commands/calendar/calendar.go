package calendar

import (
	"github.com/spf13/cobra"
)

// NewNoteCmd crea el comando padre 'note' y adjunta todos los subcomandos de notas.
func NewCalendarCmd() *cobra.Command {
	var calendarCmd = &cobra.Command{
		Use:   "calendar",
		Short: "Gestiona tu calendario",
		Long:  "Permite crear, listar, actualizar y eliminar eventos de Google Calendar.",
	}

	// Adjuntar subcomandos, pasando las dependencias a cada uno
	calendarCmd.AddCommand(NewAddCmd())
	// calendarCmd.AddCommand(NewGetCmd())
	calendarCmd.AddCommand(NewListCmd())
	calendarCmd.AddCommand(NewUpdateCmd())
	calendarCmd.AddCommand(NewDeleteCmd())
	calendarCmd.AddCommand(NewUpdateListCmd())
	calendarCmd.AddCommand(NewLoginCmd())
	calendarCmd.AddCommand(NewLogoutCmd())

	return calendarCmd
}
