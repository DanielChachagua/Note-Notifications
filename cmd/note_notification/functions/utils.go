package functions

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

func tokenFile() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, ".credentials", "calendar-go.json")
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getFreePortInRange(minPort, maxPort int) (int, error) {
	for port := minPort; port <= maxPort; port++ {
		addr := fmt.Sprintf("localhost:%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			listener.Close() 
			return port, nil 
		}
	}
	return 0, fmt.Errorf("no se encontró puerto libre entre %d y %d", minPort, maxPort)
}

func saveToken(path string, token *oauth2.Token) {
	os.MkdirAll(filepath.Dir(path), 0700)
	f, _ := os.Create(path)
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return fmt.Errorf("sistema operativo no soportado")
	}

	return cmd.Start()
}

func eventsFile() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, ".events-calendar", "calendar-events.json")
}

func saveEvents(events []*calendar.Event) error {
	path := eventsFile()

	// Crear directorio si no existe
	err := os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return fmt.Errorf("no se pudo crear directorio: %w", err)
	}

	// Crear (y sobrescribir) archivo
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("no se pudo crear archivo: %w", err)
	}
	defer f.Close()

	// Escribir JSON
	err = json.NewEncoder(f).Encode(events)
	if err != nil {
		return fmt.Errorf("error al escribir JSON: %w", err)
	}

	return nil
}