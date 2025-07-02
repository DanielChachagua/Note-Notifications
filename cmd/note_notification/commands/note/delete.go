package note

import (
	"fmt"
	"note_notifications/internal/services"

	"github.com/spf13/cobra"
)

func NewDeleteCmd(note *services.NoteService) *cobra.Command {
	var (
		id          string
	)

	cmd := &cobra.Command{
		Use:   "rm",
		Short: "eliminar una nota",
		Long:  "Eliminar una nota pro el ID.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if id == ""{
				fmt.Println("Error: Campo --id es requerido.")
				cmd.Help()
				return
			}

			err := note.Delete(id)
			if err != nil {
				fmt.Printf("Error al eliminar la nota: %v\n", err)
				return
			}

			fmt.Printf("¡Nota eliminada con éxito! ID: %s\n", id)
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "ID de la nota (requerido)")

	cmd.MarkFlagRequired("id")

	return cmd
}