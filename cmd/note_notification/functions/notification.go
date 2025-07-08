package functions

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
	"google.golang.org/api/calendar/v3"
)

type ScheduledEvent struct {
	Event  *calendar.Event
	Cancel context.CancelFunc
}

var (
	scheduledEvents = make(map[string]*ScheduledEvent)
	mu              sync.Mutex
)

func InitNotifications() error {
	// Inicializa las notificaciones y programa eventos actuales
	events, err := GetEvents()
	if err != nil {
		return err
	}

	if len(*events) == 0 {
		fmt.Println("No hay eventos próximos.")
	} else {
		var nextEvents []string
		for _, item := range *events {
			// Programa evento con contexto cancelable
			ctx, cancel := context.WithCancel(context.Background())
			scheduleEventWithContext(item, ctx)

			mu.Lock()
			scheduledEvents[item.Id] = &ScheduledEvent{Event: item, Cancel: cancel}
			mu.Unlock()

			var dateF string
			if item.Start.DateTime == "" {
				date, err := time.Parse("2006-01-02", item.Start.Date)
				if err != nil {
					fmt.Printf("Error al parsear la fecha: %v\n", err)
					return err
				}
				dateF = date.Format("02/01/2006")
			} else {
				date, err := time.Parse(time.RFC3339, item.Start.DateTime)
				if err != nil {
					fmt.Printf("Error al parsear la fecha: %v\n", err)
					return err
				}
				dateF = date.Format("02/01/2006 15:04")
			}

			nextEvents = append(nextEvents, fmt.Sprintf("%v --- %v", dateF, item.Summary))
		}

		err := listNotification(nextEvents)
		if err != nil {
			return err
		}
	}

	return nil
}

// Refresca periódicamente los eventos, agregando, actualizando o cancelando según corresponda
func RefreshEvents() error {
	mu.Lock()
	defer mu.Unlock()

	events, err := GetEvents()
	if err != nil {
		return err
	}

	newIDs := make(map[string]*calendar.Event)
	for _, e := range *events {
		newIDs[e.Id] = e
	}

	// Cancelar eventos borrados
	for id, scheduled := range scheduledEvents {
		if _, exists := newIDs[id]; !exists {
			scheduled.Cancel()
			delete(scheduledEvents, id)
			fmt.Printf("Evento borrado y cancelado: %s\n", id)
		}
	}

	// Agregar o actualizar eventos nuevos o modificados
	for id, newEvent := range newIDs {
		scheduled, exists := scheduledEvents[id]

		if !exists {
			// Evento nuevo
			ctx, cancel := context.WithCancel(context.Background())
			scheduleEventWithContext(newEvent, ctx)
			scheduledEvents[id] = &ScheduledEvent{Event: newEvent, Cancel: cancel}
			fmt.Printf("Evento nuevo programado: %s\n", id)
		} else if eventChanged(scheduled.Event, newEvent) {
			// Evento modificado
			scheduled.Cancel()
			ctx, cancel := context.WithCancel(context.Background())
			scheduleEventWithContext(newEvent, ctx)
			scheduledEvents[id] = &ScheduledEvent{Event: newEvent, Cancel: cancel}
			fmt.Printf("Evento modificado reprogramado: %s\n", id)
		} else {
			fmt.Printf("Evento sin cambios: %s\n", id)
		}
	}

	return nil
}

func scheduleEventWithContext(event *calendar.Event, ctx context.Context) {
	now := time.Now()
	eventTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
	if err != nil {
		fmt.Printf("Error al parsear la fecha: %v\n", err)
		return
	}

	before := eventTime.Add(-1 * time.Hour)
	if before.After(now) && eventTime.Before(now.Add(24*time.Hour)) {
		dur := time.Until(before)
		go func(e *calendar.Event, ctx context.Context) {
			select {
			case <-ctx.Done():
				fmt.Printf("⛔ Cancelada notificación 1 hora antes para %s\n", e.Summary)
				return
			case <-time.After(dur):
				notify(e, "Evento próximo en 1 Hora")
			}
		}(event, ctx)
	}

	durationUntil := time.Until(eventTime)
	if durationUntil > 0 && eventTime.Before(now.Add(24*time.Hour)) {
		go func(e *calendar.Event, ctx context.Context) {
			select {
			case <-ctx.Done():
				fmt.Printf("⛔ Cancelada notificación al inicio para %s\n", e.Summary)
				return
			case <-time.After(durationUntil):
				notify(e, "Comenzó un evento")
			}
		}(event, ctx)
	}
}

func eventChanged(oldEvent, newEvent *calendar.Event) bool {
	return oldEvent.Start.DateTime != newEvent.Start.DateTime ||
		oldEvent.Summary != newEvent.Summary ||
		oldEvent.Location != newEvent.Location
}

func listNotification(events []string) error {
	_, _, err := dlgs.List("Eventos Google Calendar", "Próximos Eventos:", events)
	if err != nil {
		fmt.Printf("Error al mostrar los eventos: %v", err)
		return err
	}

	return nil
}

func notify(event *calendar.Event, title string) {
	err := beeep.Notify(
		title,
		fmt.Sprintf("*Evento: %s\n   *Lugar: %s\n   *Descripción: %s", event.Summary, event.Location, event.Description),
		"",
	)

	if err != nil {
		fmt.Printf("Error al mostrar la notificación: %v\n", err)
		return
	}
}


// package functions

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"

// 	"github.com/gen2brain/beeep"
// 	"github.com/gen2brain/dlgs"
// 	"google.golang.org/api/calendar/v3"
// )

// type EventTask struct {
// 	Cancel context.CancelFunc
// }

// var (
// 	activeTasks = make(map[string]*EventTask)
// 	taskMu      sync.Mutex
// )

// type ScheduledEvent struct {
// 	Event   *calendar.Event
// 	Cancel  context.CancelFunc
// }

// var (
// 	scheduledEvents = make(map[string]*ScheduledEvent)
// 	mu             sync.Mutex
// )



// func InitNotifications() error {
// 	events, err := GetEvents()
// 	if err != nil {
// 		return err
// 	}

// 	if len(*events) == 0 {
// 		fmt.Println("No hay eventos próximos.")
// 	} else {
// 		var nextEvents []string
// 		for _, item := range *events {
// 			scheduleEvent(item)

// 			var dateF string 
// 			if item.Start.DateTime == "" {
// 				date, err := time.Parse("2006-01-02", item.Start.Date)
// 				if err != nil {
// 					fmt.Printf("Error al parsear la fecha: %v\n", err)
// 					return err
// 				}
// 				dateF = date.Format("02/01/2006")
// 			} else {
// 				date, err := time.Parse(time.RFC3339, item.Start.DateTime)
// 				if err != nil {
// 					fmt.Printf("Error al parsear la fecha: %v\n", err)
// 					return err
// 				}
// 				dateF = date.Format("02/01/2006 15:04")
// 			}

// 			nextEvents = append(nextEvents, fmt.Sprintf("%v --- %v", dateF, item.Summary))
// 		}

// 		err := listNotification(nextEvents)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func scheduleEvent(event *calendar.Event) {
// 	now := time.Now()
// 	var eventTime time.Time
// 	var err error

// 	if event.Start.DateTime == "" {
// 		return
// 	}

// 	eventTime, err = time.Parse(time.RFC3339, event.Start.DateTime)
// 	if err != nil {
// 		fmt.Printf("Error al parsear la fecha: %v\n", err)
// 		return
// 	}

// 	// Crear contexto cancelable
// 	ctx, cancel := context.WithCancel(context.Background())

// 	// Guardar el task activo
// 	taskMu.Lock()
// 	activeTasks[event.Id] = &EventTask{Cancel: cancel}
// 	taskMu.Unlock()

// 	// Notificación 1 hora antes
// 	before := eventTime.Add(-1 * time.Hour)
// 	if before.After(now) && eventTime.Before(now.Add(24*time.Hour)) {
// 		dur := time.Until(before)
// 		go func(e *calendar.Event, ctx context.Context) {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Printf("⛔ Cancelada notificación 1 hora antes para %s\n", e.Summary)
// 				return
// 			case <-time.After(dur):
// 				notify(e, "Evento próximo en 1 Hora")
// 			}
// 		}(event, ctx)
// 	}

// 	// Notificación al inicio
// 	durationUntil := time.Until(eventTime)
// 	if durationUntil > 0 && eventTime.Before(now.Add(24*time.Hour)) {
// 		go func(e *calendar.Event, ctx context.Context) {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Printf("⛔ Cancelada notificación al inicio para %s\n", e.Summary)
// 				return
// 			case <-time.After(durationUntil):
// 				notify(e, "Comenzó un evento")
// 			}
// 		}(event, ctx)
// 	}
// }

// func cancelScheduledEvent(eventID string) {
// 	taskMu.Lock()
// 	defer taskMu.Unlock()

// 	if task, ok := activeTasks[eventID]; ok {
// 		task.Cancel()
// 		delete(activeTasks, eventID)
// 		fmt.Printf("✅ Evento %s cancelado correctamente\n", eventID)
// 	} else {
// 		fmt.Printf("⚠️ Evento %s no encontrado para cancelar\n", eventID)
// 	}
// }


// // func scheduleEvent(event *calendar.Event) {
// // 	now := time.Now()
// // 	var eventTime time.Time
// // 	var err error

// // 	if event.Start.DateTime != ""{
// // 		eventTime, err = time.Parse(time.RFC3339, event.Start.DateTime)
// // 		if err != nil {
// // 			fmt.Printf("Error al parsear la fecha: %v\n", err)
// // 			return
// // 		}

// // 		before := eventTime.Add(-1 * time.Hour)
// // 		if before.After(now) && eventTime.Before(now.Add(24 * time.Hour)) {
// // 			dur := time.Until(before)
// // 			go func(e *calendar.Event) {
// // 				time.Sleep(dur)
// // 				notify(e, "Evento próximo en 1 Hora")
// // 			}(event)
// // 		}

// // 		durationUntil := time.Until(eventTime)
// // 		if durationUntil <= 0 {
// // 			return
// // 		}

// // 		if eventTime.Before(now.Add(24 * time.Hour)) {
// // 			go func(e *calendar.Event) {
// // 				time.Sleep(durationUntil)
// // 				notify(e, "Comenzó un evento")
// // 			}(event)
// // 		}
// // 	}
// // }

// func listNotification(events []string) error {
// 	_, _, err := dlgs.List("Eventos Google Calendar", "Próximos Eventos:", events)
// 	if err != nil {
// 		fmt.Printf("Error al mostrar los eventos: %v", err)
// 		return err
// 	}

// 	return nil
// }

// func notify(event *calendar.Event, title string) {
// 	err := beeep.Notify(
// 		title, 
// 		fmt.Sprintf("*Evento: %s\n   *Lugar: %s\n   *Descripción: %s", event.Summary, event.Location, event.Description), 
// 		"",
// 	)

// 	if err != nil {
// 		fmt.Printf("Error al mostrar la notificación: %v\n", err)
// 		return
// 	}
// }

// func RefreshEvents() error {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	events, err := GetEvents()
// 	if err != nil {
// 		return err
// 	}

// 	// Crear un set/map de los IDs de los eventos nuevos
// 	newIDs := make(map[string]*calendar.Event)
// 	for _, e := range *events {
// 		newIDs[e.Id] = e
// 	}

// 	// 2.a Cancelar los eventos que ya no existen
// 	for id, scheduled := range scheduledEvents {
// 		if _, exists := newIDs[id]; !exists {
// 			// Evento borrado
// 			scheduled.Cancel()
// 			delete(scheduledEvents, id)
// 			fmt.Printf("Evento borrado y cancelado: %s\n", id)
// 		}
// 	}

// 	// 2.b Agregar o actualizar eventos nuevos o modificados
// 	for id, newEvent := range newIDs {
// 		scheduled, exists := scheduledEvents[id]

// 		if !exists {
// 			// Evento nuevo: programar
// 			ctx, cancel := context.WithCancel(context.Background())
// 			scheduleEventWithContext(newEvent, ctx)
// 			scheduledEvents[id] = &ScheduledEvent{Event: newEvent, Cancel: cancel}
// 			fmt.Printf("Evento nuevo programado: %s\n", id)
// 		} else {
// 			// Evento existente: ver si cambió (ejemplo: hora)
// 			if eventChanged(scheduled.Event, newEvent) {
// 				// Cancelar la tarea anterior y reprogramar
// 				scheduled.Cancel()
// 				ctx, cancel := context.WithCancel(context.Background())
// 				scheduleEventWithContext(newEvent, ctx)
// 				scheduledEvents[id] = &ScheduledEvent{Event: newEvent, Cancel: cancel}
// 				fmt.Printf("Evento modificado reprogramado: %s\n", id)
// 			} else {
// 				fmt.Printf("Evento sin cambios: %s\n", id)
// 			}
// 		}
// 	}

// 	return nil
// }

// func scheduleEventWithContext(event *calendar.Event, ctx context.Context) {
// 	// similar a tu scheduleEvent pero usa el ctx recibido
// 	now := time.Now()
// 	eventTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
// 	if err != nil {
// 		fmt.Printf("Error al parsear la fecha: %v\n", err)
// 		return
// 	}

// 	before := eventTime.Add(-1 * time.Hour)
// 	if before.After(now) && eventTime.Before(now.Add(24*time.Hour)) {
// 		dur := time.Until(before)
// 		go func(e *calendar.Event, ctx context.Context) {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Printf("⛔ Cancelada notificación 1 hora antes para %s\n", e.Summary)
// 				return
// 			case <-time.After(dur):
// 				notify(e, "Evento próximo en 1 Hora")
// 			}
// 		}(event, ctx)
// 	}

// 	durationUntil := time.Until(eventTime)
// 	if durationUntil > 0 && eventTime.Before(now.Add(24*time.Hour)) {
// 		go func(e *calendar.Event, ctx context.Context) {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Printf("⛔ Cancelada notificación al inicio para %s\n", e.Summary)
// 				return
// 			case <-time.After(durationUntil):
// 				notify(e, "Comenzó un evento")
// 			}
// 		}(event, ctx)
// 	}
// }

// func eventChanged(oldEvent, newEvent *calendar.Event) bool {
// 	// Comparar lo que te importe, por ejemplo la hora de inicio o resumen
// 	return oldEvent.Start.DateTime != newEvent.Start.DateTime ||
// 		oldEvent.Summary != newEvent.Summary ||
// 		oldEvent.Location != newEvent.Location
// }
