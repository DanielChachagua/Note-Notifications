package functions

import "fmt"

func Login() error {
	token := GetToken()
	if token != nil {
		return nil
	}
	
	return fmt.Errorf("error al loguearse en Google Calendar")
}