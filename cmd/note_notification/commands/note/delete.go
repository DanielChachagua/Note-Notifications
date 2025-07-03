package note

import (
	"fmt"
	"note_notifications/internal/services"

	"github.com/spf13/cobra"
)

func NewDeleteCmd(note *services.NoteService) *cobra.Command {
	var (
		id string
	)

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Elimina una nota por su ID",
		Long:  "Elimina una nota de forma permanente utilizando su ID.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			err := note.Delete(id)
			if err != nil {
				fmt.Printf("Error al eliminar la nota: %v\n", err)
				return
			}

			fmt.Printf("¡Nota con ID '%s' eliminada con éxito!\n", id)
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "ID de la nota a eliminar (requerido)")

	cmd.MarkFlagRequired("id")

	return cmd
}
