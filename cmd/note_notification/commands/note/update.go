package note

import (
	"fmt"
	"note_notifications/internal/schemas"
	"note_notifications/internal/services"

	"github.com/spf13/cobra"
)

// NewUpdateCmd crea el comando para actualizar una nota existente.
// Recibe el contenedor de dependencias para acceder a los servicios necesarios.
func NewUpdateCmd(note *services.NoteService) *cobra.Command {
	var ( // Declarar variables para almacenar los valores de las flags
		id          string
		title       string
		description string
		url         string
		dateStr     string
		timeStr     string
		warn        string
	)

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Actualiza una nota existente",
		Long:  "Actualiza el título, descripción, URL, fecha y hora de una nota existente usando su ID.",
		Args:  cobra.NoArgs, // No esperamos argumentos posicionales
		Run: func(cmd *cobra.Command, args []string) {
			// 1. Parsear y validar datos de las flags
			tDate, err := schemas.ToCustomDate(dateStr)
			if err != nil {
				fmt.Printf("Error en la fecha: %v (esperado formato dd-mm-yyyy)\n", err)
				return
			}

			tTime, err := schemas.ToCustomTime(timeStr)
			if err != nil {
				fmt.Printf("Error en la hora: %v (esperado formato hh:mm)\n", err)
				return
			}

			noteData := schemas.NoteUpdate{
				ID:          id,
				Title:       title,
				Description: description,
				Url:         &url,
				Date:        tDate,
				Time:        tTime,
			}

			if warn != "" {
				value := warn == "true" || warn == "1" || warn == "t"
				noteData.Warn = &value
			}

			if err := noteData.Validate(); err != nil {
				fmt.Println(err)
				return
			}

			err = note.Update(&noteData)
			if err != nil {
				fmt.Printf("Error al actualizar la nota: %v\n", err)
				return
			}

			fmt.Printf("¡Nota con ID '%s' actualizada con éxito!\n", id)
		},
	}

	// Definir las flags para el comando 'update'
	cmd.Flags().StringVarP(&id, "id", "i", "", "ID de la nota a actualizar (requerido)")
	cmd.Flags().StringVarP(&title, "title", "n", "", "Nuevo título de la nota (requerido)")
	cmd.Flags().StringVarP(&description, "description", "b", "", "Nueva descripción de la nota (requerido)")
	cmd.Flags().StringVarP(&url, "url", "u", "", "Nueva URL asociada a la nota (opcional)")
	cmd.Flags().StringVarP(&dateStr, "date", "d", "", "Nueva fecha de la nota en formato dd-mm-yyyy (requerido)")
	cmd.Flags().StringVarP(&timeStr, "time", "t", "", "Nueva hora de la nota en formato hh:mm (requerido)")
	cmd.Flags().StringVarP(&warn, "warn", "w", "", "Nuevo aviso (warn) de la nota (opcional)")

	// Marcar flags como requeridas
	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("title")
	cmd.MarkFlagRequired("description")
	cmd.MarkFlagRequired("date")
	cmd.MarkFlagRequired("time")

	return cmd
}
