package note

import (
	"fmt"
	"note_notifications/internal/services"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// NewListCmd crea el comando para listar todas las notas.
// Recibe el contenedor de dependencias para acceder a los servicios necesarios.
func NewListCmd(note *services.NoteService) *cobra.Command {
	cmd := &cobra.Command{
		Use:	"list",
		Short:	"Listar todas las notas",
		Long:	"Muestra una lista de todas las notas existentes con sus detalles.",
		Args:	cobra.NoArgs, // No esperamos argumentos posicionales
		Run: func(cmd *cobra.Command, args []string) {
			listNotes, err := note.List()
			if err != nil {
				fmt.Printf("Error al listar las notas: %v\n", err)
				return
			}

			if len(*listNotes) == 0 {
				fmt.Println("No hay notas para mostrar.")
				return
			}

			table := tablewriter.NewTable(os.Stdout,
				tablewriter.WithHeader([]string{"ID", "Título", "Fecha", "Hora", "Aviso"}),
			)

			for _, n := range *listNotes {
				warnStr := fmt.Sprintf("%t", n.Warn)
				table.Append([]string{n.ID, n.Title, n.Date.ToTime().Format("02-01-2006"), n.Time.ToTime().Format("15:04"), warnStr})
			}

			table.Render()
		},
	}

	return cmd
}


