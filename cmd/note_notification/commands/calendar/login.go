package calendar

import (
	"fmt"
	"note_notifications/cmd/note_notification/functions"

	"github.com/spf13/cobra"
)

// NewListCmd crea el comando para listar todas las notas.
// Recibe el contenedor de dependencias para acceder a los servicios necesarios.
func NewLoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Loguearse en Google Calendar",
		Long:  "Loguearse en Google Calendar para poder crear, listar, actualizar y eliminar eventos de Google Calendar.",
		Args:  cobra.NoArgs, // No esperamos argumentos posicionales
		Run: func(cmd *cobra.Command, args []string) {

			err := functions.Login()
			if err != nil {
				fmt.Printf("Error al loguearse en Google Calendar: %v\n", err)
				return
			}

			fmt.Println("✅ Logueado correctamente en Google Calendar")
		},
	}

	return cmd
}