package calendar

import (
	"fmt"
	"note_notifications/cmd/note_notification/functions"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"google.golang.org/api/calendar/v3"
)

// NewListCmd crea el comando para listar todas las notas.
// Recibe el contenedor de dependencias para acceder a los servicios necesarios.
func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Listar todas las notas",
		Long:  "Muestra una lista de todas las notas existentes con sus detalles.",
		Args:  cobra.NoArgs, // No esperamos argumentos posicionales
		Run: func(cmd *cobra.Command, args []string) {

			// var events *[]*calendar.Event
			events := &[]*calendar.Event{}

			notExist, err := functions.FileDoesNotExistOrIsEmpty(functions.EventsFile())
			if err != nil {
				fmt.Printf("Error al listar las notas: %v\n", err)
				return
			}

			if notExist {
				events, err = functions.GetEvents()
				if err != nil {
					fmt.Printf("Error al listar las notas: %v\n", err)
					return
				}
			} else {
				events, err = functions.GetEventsFromFile()
				if err != nil {
					fmt.Printf("Error al listar las notas: %v\n", err)
					return
				}				
			}

			if len(*events) == 0 {
					fmt.Println("No hay eventos para mostrar.")
					return
				}

				table := tablewriter.NewTable(os.Stdout,
					tablewriter.WithHeader([]string{"ID", "Summary", "Description", "Start"}),
				)

				for _, e := range *events {
					// Usar Start.DateTime o Start.Date según el tipo de evento
					date := e.Start.DateTime
					if date == "" {
						date = e.Start.Date
					}

					table.Append([]string{
						e.Id,
						e.Summary,
						e.Description,
						date,
					})
				}

				table.Render()
		},
	}

	return cmd
}
