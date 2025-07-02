package note

// import (
// 	"fmt"
// 	"note_notifications/internal/schemas"
// 	"note_notifications/internal/services"

// 	"github.com/spf13/cobra"
// )

// // NewAddCmd crea el comando para agregar una nueva nota.
// // Recibe el contenedor de dependencias para acceder a los servicios necesarios.
// func NewAddCmd(note *services.NoteService) *cobra.Command {
// 	var ( // Declarar variables para almacenar los valores de las flags
// 		title       string
// 		description string
// 		url         string
// 		dateStr     string
// 		timeStr     string
// 	)

// 	cmd := &cobra.Command{
// 		Use:   "add",
// 		Short: "Agregar una nueva nota",
// 		Long:  "Agregar una nueva nota con título, descripción, URL, fecha y hora usando flags.",
// 		Args:  cobra.NoArgs, // No esperamos argumentos posicionales
// 		Run: func(cmd *cobra.Command, args []string) {
// 			// Validar que las flags requeridas no estén vacías
// 			if title == "" || description == "" || dateStr == "" || timeStr == "" {
// 				fmt.Println("Error: Los campos --title, --description, --date y --time son requeridos.")
// 				cmd.Help()
// 				return
// 			}

// 			// 1. Parsear y validar datos de las flags
// 			tDate, err := schemas.ToCustomDate(dateStr)
// 			if err != nil {
// 				fmt.Printf("Error en la fecha: %v (esperado formato dd-mm-yyyy)\n", err)
// 				return
// 			}

// 			tTime, err := schemas.ToCustomTime(timeStr)
// 			if err != nil {
// 				fmt.Printf("Error en la hora: %v (esperado formato hh:mm)\n", err)
// 				return
// 			}

// 			noteData := schemas.NoteCreate{
// 				Title:       title,
// 				Description: description,
// 				Url:         &url, // La URL puede ser opcional, por eso se pasa su puntero
// 				Date:        tDate,
// 				Time:        tTime,
// 			}

// 			if err := noteData.Validate(); err != nil {
// 				fmt.Println(err)
// 				return
// 			}

// 			// 2. Usar el servicio del contenedor de dependencias
// 			createdNote, err := note.Create(&noteData)
// 			if err != nil {
// 				fmt.Printf("Error al crear la nota: %v\n", err)
// 				return
// 			}

// 			fmt.Printf("¡Nota creada con éxito! ID: %s\n", createdNote)
// 		},
// 	}

// 	// Definir las flags para el comando 'add'
// 	cmd.Flags().StringVarP(&title, "title", "n", "", "Título de la nota (requerido)")
// 	cmd.Flags().StringVarP(&description, "description", "b", "", "Descripción de la nota (requerido)")
// 	cmd.Flags().StringVarP(&url, "url", "u", "", "URL asociada a la nota (opcional)")
// 	cmd.Flags().StringVarP(&dateStr, "date", "d", "", "Fecha de la nota en formato dd-mm-yyyy (requerido)")
// 	cmd.Flags().StringVarP(&timeStr, "time", "t", "", "Hora de la nota en formato hh:mm (requerido)")

// 	// Marcar flags como requeridas
// 	cmd.MarkFlagRequired("title")
// 	cmd.MarkFlagRequired("description")
// 	cmd.MarkFlagRequired("date")
// 	cmd.MarkFlagRequired("time")

// 	return cmd
// }