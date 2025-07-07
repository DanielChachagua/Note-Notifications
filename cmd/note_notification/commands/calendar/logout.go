package calendar

import (
	"fmt"
	"note_notifications/cmd/note_notification/functions"

	"github.com/spf13/cobra"
)

// NewListCmd crea el comando para listar todas las notas.
// Recibe el contenedor de dependencias para acceder a los servicios necesarios.
func NewLogoutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Cerrar sesión de Google Calendar",
		Long:  "Desloguearse de Google Calendar.",
		Args:  cobra.NoArgs, // No esperamos argumentos posicionales
		Run: func(cmd *cobra.Command, args []string) {

			err := functions.Logout()
			if err != nil {
				fmt.Printf("Error al loguearse en Google Calendar: %v\n", err)
				return
			}

			fmt.Println("✅ Logout exitoso. El token fue eliminado.")
		},
	}

	return cmd
}