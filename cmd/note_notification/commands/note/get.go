package note

import (
	"fmt"
	"note_notifications/internal/services"

	"github.com/spf13/cobra"
)

func NewGetCmd(note *services.NoteService) *cobra.Command {
	var (
		id       string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Obtener una nota por su ID",
		Long:  "Obtiene los detalles de una nota específica utilizando su ID.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			retrievedNote, err := note.Get(id)
			if err != nil {
				fmt.Printf("Error al obtener la nota: %v\n", err)
				return
			}

			fmt.Printf("Nota encontrada:\n")
			fmt.Printf("  ID:          %s\n", retrievedNote.ID)
			fmt.Printf("  Título:      %s\n", retrievedNote.Title)
			fmt.Printf("  Descripción: %s\n", retrievedNote.Description)
			if retrievedNote.Url != nil && *retrievedNote.Url != "" {
				fmt.Printf("  URL:         %s\n", *retrievedNote.Url)
			}
			fmt.Printf("  Fecha:       %s\n", retrievedNote.Date.ToTime().Format("02-01-2006"))
            fmt.Printf("  Hora:        %s\n", retrievedNote.Time.ToTime().Format("15:04"))
			fmt.Printf("  Aviso:       %t\n", retrievedNote.Warn)
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "ID de la nota (requerido)")

	cmd.MarkFlagRequired("id")

	return cmd
}