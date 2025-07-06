package calendar

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	var (
		id string
	)

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Elimina una nota por su ID",
		Long:  "Elimina una nota de forma permanente utilizando su ID.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
		
			fmt.Printf("¡Nota con ID '%s' eliminada con éxito!\n", id)
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "ID de la nota a eliminar (requerido)")

	cmd.MarkFlagRequired("id")

	return cmd
}
