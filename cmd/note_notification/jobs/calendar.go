package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"

	"golang.org/x/oauth2"
)

// Lee las credenciales y realiza la autenticación
func GetClient(config *oauth2.Config) *http.Client {
	tokFile := tokenFile()
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func tokenFile() string {
	usr, _ := user.Current()
	return filepath.Join(usr.HomeDir, ".credentials", "calendar-go.json")
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	config.RedirectURL = "http://localhost:8080/callback"
	codeCh := make(chan string)
	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "No se recibió el código", http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, "Autenticación exitosa. Podés cerrar esta pestaña.")
		codeCh <- code

		// Cerramos el servidor
		go func() {
			_ = srv.Shutdown(context.Background())
		}()
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar servidor local: %v", err)
		}
	}()

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Abrí este link en tu navegador y pegá el código acá:\n%v\n", authURL)

	err := openBrowser(authURL)
	if err != nil {
		fmt.Println("Error al abrir el navegador:", err)
	}

	code := <-codeCh

	tok, err := config.Exchange(context.TODO(), code)
	if err != nil {
		log.Fatalf("No se pudo obtener el token: %v", err)
	}
	return tok
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
