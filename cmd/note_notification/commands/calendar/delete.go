package calendar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"note_notifications/cmd/note_notification/functions"
	"note_notifications/internal/schemas"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	var ids []string

	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Elimina una o más notas por sus IDs",
		Long:  "Elimina notas permanentemente pasando uno o más IDs separados por espacio después del flag -i.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if len(ids) == 0 {
				fmt.Println("❌ Debes usar -i seguido de al menos un ID")
				return
			}

			token := functions.GetToken()

			calendarDelete := schemas.DeleteEvent{
				Token:    token,
				EventIds: ids,
			}

			jsonData, err := json.Marshal(calendarDelete)
			if err != nil {
				fmt.Printf("Error al serializar IDs: %v\n", err)
				return
			}

			response, err := http.Post("http://localhost:3000/calendar/delete", "application/json", bytes.NewReader([]byte(jsonData)))
			if err != nil {
				fmt.Printf("Error al eliminar las notas: %v\n", err)
				return
			}
			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("❌ Error al leer el cuerpo de la respuesta: %v\n", err)
				return
			}

			if response.StatusCode != http.StatusOK {
				fmt.Printf("❌ Error del servidor (HTTP %d): %s\n", response.StatusCode, string(body))
				return
			}

			err = functions.DeleteEventFronJson(ids)
			if err != nil {
				fmt.Printf("Error al eliminar los eventos de forma local: %v\n", err)
				fmt.Print("\nActualice manualmente el la memoria local con el comando 'ntn calendar update'")
			}

			fmt.Printf("¡Notas con IDs '%s' eliminadas con éxito!\n", ids)
		},
	}

	cmd.Flags().StringSliceVarP(&ids, "id", "i", nil, "IDs de las notas a eliminar (requerido)")

	cmd.MarkFlagRequired("id")

	return cmd
}
