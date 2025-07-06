package calendar

import (
	"github.com/spf13/cobra"
)

func NewGetCmd() *cobra.Command {
	var (
		id       string
	)

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Obtener una nota por su ID",
		Long:  "Obtiene los detalles de una nota específica utilizando su ID.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {

			
			
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "ID de la nota (requerido)")

	cmd.MarkFlagRequired("id")

	return cmd
}