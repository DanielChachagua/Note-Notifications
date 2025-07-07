package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"google.golang.org/api/calendar/v3"
)

func AddEvent(event *calendar.Event) error {
	jsonEvents, err := os.ReadFile(EventsFile())
	if err != nil {
		return fmt.Errorf("error al leer archivo: %w", err)
	}

	var events []*calendar.Event
	err = json.Unmarshal([]byte(jsonEvents), &events)
	if err != nil {
		return fmt.Errorf("error al obtener eventos: %w", err)
	}

	events = append(events, event)

	sort.Slice(events, func(i, j int) bool {
		startI := parseStartTime(events[i])
		startJ := parseStartTime(events[j])
		return startI.Before(startJ)
	})

	return saveEvents(events)
}