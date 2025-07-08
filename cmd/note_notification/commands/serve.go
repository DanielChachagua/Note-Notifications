package commands

import (
	"fmt"
	"note_notifications/cmd/note_notification/functions"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Inicia el servicio de notificaciones en segundo plano",
	Long:  `Inicia el servicio que monitorea y envía notificaciones para los eventos próximos. Este comando debe mantenerse en ejecución para que las notificaciones funcionen.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Iniciando el servicio de notificaciones...")

		// Inicia las notificaciones en una goroutine para no bloquear
		// go func() {
			if err := functions.InitNotifications(); err != nil {
				fmt.Printf("Error al iniciar las notificaciones: %v\n", err)
			}
		// }()
		fmt.Println("Servicio de notificaciones corriendo. Presiona Ctrl+C para detener.")

		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		for {
			select {
			case <-ticker.C:
				if err := functions.RefreshEvents(); err != nil {
					fmt.Printf("Error refrescando eventos: %v\n", err)
				}
			case <-quit:
				fmt.Println("Terminando servicio de notificaciones")
				return nil
			}
		}

		// Espera una señal de interrupción para terminar de forma limpia
		// quit := make(chan os.Signal, 1)
		// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		// <-quit

		// fmt.Println("Deteniendo el servicio de notificaciones...")
	},
}

// func init() {
// 	rootCmd.AddCommand(serveCmd)
// }



// package commands

// import (
// 	"fmt"
// 	"note_notifications/cmd/note_notification/functions"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/spf13/cobra"
// )

// var serveCmd = &cobra.Command{
// 	Use:	"serve",
// 	Short:	"Inicia el servicio de notificaciones en segundo plano",
// 	Long:	`Inicia el servicio que monitorea y envía notificaciones para los eventos próximos. Este comando debe mantenerse en ejecución para que las notificaciones funcionen.`,
// 	RunE:	func(cmd *cobra.Command, args []string) error {
// 		fmt.Println("Iniciando el servicio de notificaciones...")

// 		// Inicia las notificaciones en una goroutine para no bloquear
// 		go func() {
// 			if err := functions.InitNotifications(); err != nil {
// 				fmt.Printf("Error al iniciar las notificaciones: %v\n", err)
// 			}
// 		}()

// 		fmt.Println("Servicio de notificaciones corriendo. Presiona Ctrl+C para detener.")

// 		// Espera una señal de interrupción para terminar de forma limpia
// 		quit := make(chan os.Signal, 1)
// 		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 		<-quit

// 		fmt.Println("Deteniendo el servicio de notificaciones...")
// 		return nil
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(serveCmd)
// }