package calendar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	cal "google.golang.org/api/calendar/v3"
	"io"
	"net/http"
	"note_notifications/cmd/note_notification/functions"
	"note_notifications/internal/schemas"
)

// NewUpdateCmd crea el comando para actualizar una nota existente.
// Recibe el contenedor de dependencias para acceder a los servicios necesarios.
func NewUpdateCmd() *cobra.Command {
	var ( // Declarar variables para almacenar los valores de las flags
		id          string
		summary     string
		location    string
		description string
		date        string
		time        string
	)

	cmd := &cobra.Command{
		Use:   "put",
		Short: "Actualiza una nota existente",
		Long:  "Actualiza el título, descripción, URL, fecha y hora de una nota existente usando su ID.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			tDate, err := schemas.ToCustomDate(date)
			if err != nil {
				fmt.Printf("Error en la fecha: %v (esperado formato dd-mm-yyyy)\n", err)
				return
			}

			var tTime *schemas.CustomTime
			if cmd.Flags().Changed("time") {
				t, err := schemas.ToCustomTime(time)
				if err != nil {
					fmt.Printf("❌ Error en la hora: %v (esperado formato hh:mm)\n", err)
					return
				}
				tTime = &t
			}

			var locationPtr *string
			if cmd.Flags().Changed("location") {
				locationPtr = &location
			}

			var descriptionPtr *string
			if cmd.Flags().Changed("description") {
				descriptionPtr = &description
			}

			calendar := schemas.CalendarUpdate{
				ID:          id,
				Summary:     summary,
				Location:    locationPtr,
				Description: descriptionPtr,
				Date:        tDate,
				Time:        tTime,
			}

			if err := calendar.Validate(); err != nil {
				fmt.Println(err)
				return
			}

			token := functions.GetToken()

			updateEvent := schemas.UpdateEvent{
				Token: token,
				Event: calendar,
			}

			if err := updateEvent.Validate(); err != nil {
				fmt.Println(err)
				return
			}

			calendarData, err := json.Marshal(updateEvent)
			if err != nil {
				fmt.Printf("Error al codificar la nota: %v\n", err)
				return
			}

			response, err := http.Post("https://calendar.saltaget.com/calendar/update", "application/json", bytes.NewReader([]byte(calendarData)))
			if err != nil {
				fmt.Printf("Error al actualizar la nota: %v\n", err)
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

			var resp struct {
				Updated *cal.Event `json:"updated"`
			}

			err = json.Unmarshal(body, &resp)
			if err != nil {
				fmt.Printf("Error al decodificar JSON: %v", err)
			}

			err = functions.UpdateEventInJson(resp.Updated)
			if err != nil {
				fmt.Printf("Error al actualizar el evento localmente: %v", err)
				fmt.Print("\nActualice manualmente el la memoria local con el comando 'ntn calendar update'")
			}

			fmt.Printf("Evento con ID '%s' actualizada con éxito!\n", id)
		},
	}

	// Definir las flags para el comando 'update'
	cmd.Flags().StringVarP(&id, "id", "i", "", "ID de la nota a actualizar (requerido)")
	cmd.Flags().StringVarP(&summary, "summary", "s", "", "Título de la nota (requerido)")
	cmd.Flags().StringVarP(&location, "location", "l", "", "URL asociada a la nota (opcional)")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Descripción de la nota (opcional)")
	cmd.Flags().StringVarP(&date, "date", "D", "", "Fecha de la nota en formato dd-mm-yyyy (requerido)")
	cmd.Flags().StringVarP(&time, "time", "T", "", "Hora de la nota en formato hh:mm (opcional)")

	// Marcar flags como requeridas
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("summary")
	cmd.MarkFlagRequired("date")

	return cmd
}
