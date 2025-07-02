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
		Short: "Agregar una nueva nota",
		Long:  "Agregar una nueva nota con título, descripción, URL, fecha y hora usando flags.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if id == "" {
				fmt.Println("Error: Campo --id es requerido.")
				cmd.Help()
				return
			}

			createdNote, err := note.Get(id)
			if err != nil {
				fmt.Printf("Error al obtener la nota: %v\n", err)
				return
			}

			fmt.Printf("¡Nota obtenida con éxito! ID: %v\n", createdNote)
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "ID de la nota (requerido)")

	cmd.MarkFlagRequired("id")

	return cmd
}