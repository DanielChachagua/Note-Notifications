package calendar

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// NewListCmd crea el comando para listar todas las notas.
// Recibe el contenedor de dependencias para acceder a los servicios necesarios.
func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:	"list",
		Short:	"Listar todas las notas",
		Long:	"Muestra una lista de todas las notas existentes con sus detalles.",
		Args:	cobra.NoArgs, // No esperamos argumentos posicionales
		Run: func(cmd *cobra.Command, args []string) {

			table := tablewriter.NewTable(os.Stdout,
				tablewriter.WithHeader([]string{"ID", "Título", "Fecha", "Hora", "Aviso"}),
			)

			table.Render()
		},
	}

	return cmd
}


