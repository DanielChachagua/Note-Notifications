package functions

import "fmt"

func InitNotifications() error {
	events, err := GetEvents()
	if err != nil {
		return err
	}

	if len(*events) == 0 {
		fmt.Println("No hay eventos próximos.")
	} else {
		for _, item := range *events {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v - %s\n", date, item.Summary)
			
			fmt.Printf("%+v\n", item)
		}
	}

	return nil
}
