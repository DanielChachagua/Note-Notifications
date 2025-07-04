package note

import (
	"note_notifications/internal/services"

	"github.com/spf13/cobra"
)

// NewNoteCmd crea el comando padre 'note' y adjunta todos los subcomandos de notas.
func NewNoteCmd(note *services.NoteService) *cobra.Command {
	var noteCmd = &cobra.Command{
		Use:   "note",
		Short: "Gestiona tus notas",
		Long:  "Permite crear, listar, actualizar y eliminar notas.",
	}

	// Adjuntar subcomandos, pasando las dependencias a cada uno
	noteCmd.AddCommand(NewAddCmd(note))
	noteCmd.AddCommand(NewGetCmd(note))
	noteCmd.AddCommand(NewListCmd(note))
	noteCmd.AddCommand(NewUpdateCmd(note))
	noteCmd.AddCommand(NewDeleteCmd(note))

	return noteCmd
}
