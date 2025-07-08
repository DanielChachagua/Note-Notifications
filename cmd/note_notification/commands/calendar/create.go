package calendar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	cal "google.golang.org/api/calendar/v3"
	"net/http"
	"note_notifications/cmd/note_notification/functions"
	"note_notifications/internal/schemas"
)

// NewAddCmd crea el comando para agregar una nueva nota.
// Recibe el contenedor de dependencias para acceder a los servicios necesarios.
func NewAddCmd() *cobra.Command {
	var ( // Declarar variables para almacenar los valores de las flags
		summary     string
		location    string
		description string
		date        string
		time        string
	)

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Agregar una nueva nota",
		Long:  "Agregar una nueva nota con título, descripción, URL, fecha y hora usando flags.",
		Args:  cobra.NoArgs, // No esperamos argumentos posicionales
		Run: func(cmd *cobra.Command, args []string) {

			tDate, err := schemas.ToCustomDate(date)
			if err != nil {
				fmt.Printf("Error en la fecha: %v\n", err)
				return
			}

			var tTime schemas.CustomTime
			if time != "" {
				tTime, err = schemas.ToCustomTime(time)
				if err != nil {
					fmt.Printf("Error en la hora: %v\n", err)
					return
				}
			}

			calendar := schemas.CalendarCreate{
				Summary:     summary,
				Location:    &location,
				Description: &description,
				Date:        tDate,
			}

			if tTime.ToTime().IsZero() {
				calendar.Time = nil
			} else {
				calendar.Time = &tTime
			}

			if err := calendar.Validate(); err != nil {
				fmt.Println(err)
				return
			}

			token := functions.GetToken()

			calendarCreate := schemas.CreateEvent{
				Token: token,
				Event: calendar,
			}

			var created struct {
				Event *cal.Event `json:"event"`
			}

			calendarData, err := json.Marshal(calendarCreate)
			if err != nil {
				fmt.Printf("error al codificar el token: %v", err)
			} else {
				response, err := http.Post("https://calendar.saltaget.com/calendar/create", "application/json", bytes.NewReader(calendarData))
				if err != nil {
					fmt.Printf("error al obtener los eventos: %v", err)
				}
				defer response.Body.Close()

				err = json.NewDecoder(response.Body).Decode(&created)
				if err != nil {
					fmt.Printf("Error al decodificar JSON: %v", err)
				}

				err = functions.AddEvent(created.Event)
				if err != nil {
					fmt.Printf("Error al agregar el evento localmente: %v", err)
					fmt.Print("\nActualice manualmente el la memoria local con el comando 'ntn calendar update'")
				}

				fmt.Println("Evento creado exitosamente")
			}
		},
	}

	// Definir las flags para el comando 'add'
	cmd.Flags().StringVarP(&summary, "summary", "s", "", "Título de la nota (requerido)")
	cmd.Flags().StringVarP(&location, "location", "l", "", "URL asociada a la nota (opcional)")
	cmd.Flags().StringVarP(&description, "description", "d", "", "Descripción de la nota (opcional)")
	cmd.Flags().StringVarP(&date, "date", "D", "", "Fecha de la nota en formato dd-mm-yyyy (requerido)")
	cmd.Flags().StringVarP(&time, "time", "T", "", "Hora de la nota en formato hh:mm (opcional)")

	// Marcar flags como requeridas
	cmd.MarkFlagRequired("summary")
	cmd.MarkFlagRequired("date")

	return cmd
}
