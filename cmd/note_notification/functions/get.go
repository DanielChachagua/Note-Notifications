package functions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

func GetToken() *oauth2.Token {
	tokFile := TokenFile()
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = GetTokenFromWeb()
		if tok == nil {
			log.Fatal("Error al obtener token, cerrando aplicación...")
			return nil
		}
		saveToken(tokFile, tok)
	}
	return tok
}

func GetTokenFromWeb() *oauth2.Token {
	port, err := getFreePortInRange(9000, 9900)
	if err != nil {
		log.Fatalf("Error al buscar puerto libre: %v", err)
	}

	redirectURL := fmt.Sprintf("http://localhost:%d/callback", port)
	log.Println("Usando puerto dinámico")

	codeCh := make(chan string)
	srv := &http.Server{Addr: fmt.Sprintf(":%d", port)}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "No se recibió el código", http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, "Autenticación exitosa. Podés cerrar esta pestaña.")
		codeCh <- code

		go func() {
			_ = srv.Shutdown(context.Background())
		}()
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar servidor local: %v", err)
		}
	}()

	resp, err := http.Get("https://calendar.saltaget.com/calendar/get_url?redirect_url=" + redirectURL)
	if err != nil {
		log.Fatalf("No se pudo obtener el token: %v", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("No se pudo obtener el token: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		URL string `json:"url"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Error al decodificar JSON: %v", err)
	}

	fmt.Println("URL recibida:", result.URL)

	err = openBrowser(result.URL)
	if err != nil {
		fmt.Println("Error al abrir el navegador:", err)
	}

	// code := <-codeCh
	var code string
	select {
	case code = <-codeCh:
		log.Println("Código recibido")
	case <-time.After(60 * time.Second):
		fmt.Println("⏱️ Tiempo de espera agotado. No se recibió respuesta de autenticación.")
		_ = srv.Shutdown(context.Background())
		return nil
	}

	log.Println("Code:", code)

	token, err := http.Post(fmt.Sprintf("https://calendar.saltaget.com/calendar/get_token?code=%s&redirect_url=%s", code, redirectURL), "", nil)
	if err != nil {
		log.Fatalf("No se pudo obtener el token: %v", err)
	}
	defer token.Body.Close()

	fmt.Println("Token:", token)

	var tok struct {
		Token *oauth2.Token `json:"token"`
	}

	err = json.NewDecoder(token.Body).Decode(&tok)
	if err != nil {
		log.Fatalf("Error al decodificar JSON: %v", err)
	}

	fmt.Println("URL recibida:", tok.Token)

	return tok.Token
}

func GetEvents() (*[]*calendar.Event, error) {
	token := GetToken()

	tokenJson, err := json.Marshal(token)
	if err != nil {
		return nil, fmt.Errorf("error al codificar el token: %v", err)
	}

	events, err := http.Post("https://calendar.saltaget.com/calendar/get_events", "application/json", bytes.NewReader(tokenJson))
	if err != nil {
		return nil, fmt.Errorf("error al obtener los eventos: %v", err)
	}
	defer events.Body.Close()

	var result struct {
		Items []*calendar.Event `json:"items"`
	}

	if err := json.NewDecoder(events.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta: %v", err)
	}

	sort.Slice(result.Items, func(i, j int) bool {
		startI := parseStartTime(result.Items[i])
		startJ := parseStartTime(result.Items[j])
		return startI.Before(startJ)
	})

	err = saveEvents(result.Items)
	if err != nil {
		return nil, fmt.Errorf("error al guardar los eventos: %v", err)
	}

	return &result.Items, nil
}

func GetEventsFromFile() (*[]*calendar.Event, error) {
	jsonEvents, err := os.ReadFile(EventsFile())
	if err != nil {
		return nil, fmt.Errorf("error al leer archivo: %w", err)
	}

	var events []*calendar.Event
	err = json.Unmarshal([]byte(jsonEvents), &events)
	if err != nil {
		return nil, fmt.Errorf("error al obtener eventos: %w", err)
	}

	return &events, nil
}
