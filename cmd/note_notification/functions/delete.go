package functions

import (
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/api/calendar/v3"
)

func DeleteEventFronJson(eventIds []string) error {
	jsonEvents, err := os.ReadFile(EventsFile())
	if err != nil {
		return fmt.Errorf("error al leer archivo: %w", err)
	}

	var events []*calendar.Event
	err = json.Unmarshal([]byte(jsonEvents), &events)
	if err != nil {
		return fmt.Errorf("error al obtener eventos: %w", err)
	}

	toDelete := make(map[string]struct{}, len(eventIds))
	for _, id := range eventIds {
		toDelete[id] = struct{}{}
	}

	if len(toDelete) == 0 {
		return fmt.Errorf("no se encontraron IDs para eliminar")
	}

	filteredEvents := events[:0] 
	for _, ev := range events {
		if _, found := toDelete[ev.Id]; !found {
			filteredEvents = append(filteredEvents, ev)
		}
	}

	newData, err := json.MarshalIndent(filteredEvents, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar eventos: %w", err)
	}

	err = os.WriteFile(EventsFile(), newData, 0644)
	if err != nil {
		return fmt.Errorf("error al guardar archivo: %w", err)
	}

	return nil
}