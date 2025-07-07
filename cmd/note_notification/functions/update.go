package functions

import (
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/api/calendar/v3"
)

func UpdateEventInJson(updatedEvent *calendar.Event) error {
	// Leer archivo
	jsonEvents, err := os.ReadFile(EventsFile())
	if err != nil {
		return fmt.Errorf("error al leer archivo: %w", err)
	}

	// Deserializar eventos
	var events []*calendar.Event
	err = json.Unmarshal(jsonEvents, &events)
	if err != nil {
		return fmt.Errorf("error al deserializar eventos: %w", err)
	}

	// Buscar y actualizar el evento correspondiente
	found := false
	for i, ev := range events {
		if ev.Id == updatedEvent.Id {
			events[i] = updatedEvent
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("evento con ID %s no encontrado", updatedEvent.Id)
	}

	// Volver a guardar
	newData, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar eventos: %w", err)
	}

	err = os.WriteFile(EventsFile(), newData, 0644)
	if err != nil {
		return fmt.Errorf("error al guardar archivo: %w", err)
	}

	return nil
}