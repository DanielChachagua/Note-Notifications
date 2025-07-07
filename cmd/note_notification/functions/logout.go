package functions

import (
	"fmt"
	"os"
)

func Logout() error {
	tokenPath := TokenFile()

	err := os.Remove(tokenPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("✅ Ya estás deslogueado. No hay token guardado.")
			return nil
		}
		return fmt.Errorf("❌ Error al intentar eliminar el token: %w", err)
	}

	eventPath := EventsFile()

	err = os.Remove(eventPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("✅ Ya estás deslogueado. No hay eventos guardados localmente.")
			return nil
		}
		return fmt.Errorf("❌ Error al intentar eliminar eventos locales: %w", err)
	}

	return nil
}