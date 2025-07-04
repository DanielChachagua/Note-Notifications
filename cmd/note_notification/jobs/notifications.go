package jobs

// import (
// 	"fmt"
// 	"net/url"
// 	"note_notifications/internal/dependencies"
// 	"note_notifications/internal/schemas"
// 	"os/exec"
// 	"runtime"
// 	"strings"
// 	"time"

// 	"github.com/gen2brain/dlgs"
// 	webview "github.com/webview/webview_go"
// )

// func InitNotifications(deps *dependencies.Container, noteChan chan []schemas.NoteResponse) {
// 	notes, err := deps.Services.Note.GetAllWarn()
// 	if err != nil {
// 		return
// 	}

// 	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
// 	if err != nil {
// 		fmt.Println(err)
// 	}


// 	// ShowAllNotesInOneWindow(notes)

// 	for _, note := range *notes {
// 		n := note // copia local para que la goroutine no comparta el puntero
// 		go scheduleNote(n, noteChan)
// 	}
// }

// func scheduleNote(note schemas.NoteResponse, noteChan chan []schemas.NoteResponse) {
// 	// Parseamos la fecha y hora
// 	layout := "02-01-2006 15:04"
// 	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
// 	datetimeStr := fmt.Sprintf("%s %s", note.Date.String(), note.Time.String())

// 	execTime, err := time.ParseInLocation(layout, datetimeStr, loc)
// 	if err != nil {
// 		fmt.Println("Error al parsear fecha y hora:", err)
// 		return
// 	}
// 	danielmchachagua@gmail.com

// 	fmt.Print(datetimeStr, " - ", execTime, " - ", time.Now(), "\n")
// 	delay := time.Until(execTime)

// 	if delay <= 0 {
// 		fmt.Println("La nota ya expiró:", note.Title)
// 		return
// 	}

// 	// Ejecutamos la nota cuando llegue el momento
// 	time.AfterFunc(delay, func() {
// 		// notes := make([]schemas.NoteResponse, 1)
// 		// notes[0] = note
// 		// showAllNotesInOneWindow(&notes)
// 		noteChan <- []schemas.NoteResponse{note}
// 	})
// }


// func ShowAllNotesInOneWindow(notes *[]schemas.NoteResponse) {
// 	if notes == nil || len(*notes) == 0 {
// 		return
// 	}

// 	var builder strings.Builder
// 	builder.WriteString(`
// 		<html>
// 		<head>
// 			<meta charset="UTF-8">
// 			<style>
// 				body {
// 					font-family: Arial, sans-serif;
// 					background-color: #f5f5f5;
// 					margin: 0;
// 					padding: 0;
// 					display: flex;
// 					flex-direction: column;
// 					height: 100vh;
// 					color: #333;
// 				}
// 				h3 {
// 					color: #007acc;
// 					text-align: center;
// 					margin: 16px 0;
// 					flex: 0 0 auto;
// 				}
// 				.grid-container {
// 					flex: 1 1 auto;
// 					overflow-y: auto;
// 					padding: 0 20px 0px 20px;
// 				}
// 				.grid {
// 					display: grid;
// 					grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
// 					gap: 15px;
// 					margin-bottom: 20px;
// 				}
// 				.note {
// 					background: #fff;
// 					border: 1px solid #ccc;
// 					border-radius: 8px;
// 					padding: 12px 16px;
// 					box-shadow: 0 2px 5px rgba(0,0,0,0.05);
// 				}
// 				.note-title {
// 					font-weight: bold;
// 					font-size: 16px;
// 					margin-bottom: 6px;
// 				}
// 				.note-meta {
// 					font-size: 13px;
// 					color: #666;
// 					margin-bottom: 6px;
// 				}
// 				.note-desc {
// 					font-size: 14px;
// 				}
// 				.actions {
// 					flex: 0 0 auto;
// 					padding: 12px 20px;
// 					background-color: #fff;
// 					border-top: 1px solid #ccc;
// 					text-align: center;
// 					position: sticky;
// 					bottom: 0;
// 					box-shadow: 0 -2px 5px rgba(0,0,0,0.05);
// 				}
// 				button {
// 					padding: 8px 16px;
// 					font-size: 14px;
// 					background-color: #007acc;
// 					color: white;
// 					border: none;
// 					border-radius: 4px;
// 					cursor: pointer;
// 				}
// 				button:hover {
// 					background-color: #005f99;
// 				}
// 				a {
// 					color: #007acc;
// 					text-decoration: none;
// 				}
// 				a:hover {
// 					text-decoration: underline;
// 				}
// 			</style>
// 			</head>
// 			<body>
// 				<h3>Notas con aviso</h3>
// 				<div class="grid-container">
// 					<div class="grid">
// 	`)

// 	for _, note := range *notes {
// 		builder.WriteString("<div class='note'>")
// 		builder.WriteString(fmt.Sprintf("<div class='note-title'>%s</div>", note.Title))
// 		builder.WriteString(fmt.Sprintf("<div class='note-meta'>📅 %s 🕒 %s</div>", note.Date.ToTime().Format("2006-01-02"), note.Time.ToTime().Format("15:04")))
// 		builder.WriteString(fmt.Sprintf("<div class='note-desc'>%s</div>", note.Description))
// 		if note.Url != nil && *note.Url != "" {
// 			builder.WriteString(fmt.Sprintf(
// 				`<br><a href="#" onclick="window.open('%s')">%s</a>`,
// 				*note.Url, *note.Url))
// 		}
// 		builder.WriteString("</div>")
// 	}

// 	builder.WriteString(`</div>`)
// 	builder.WriteString(`<div class="actions"><button onclick="window.close()">Aceptar</button></div>`)
// 	builder.WriteString(`</body></html>`)

// 	htmlContent := builder.String()

// 	debug := true
// 	w := webview.New(debug)
// 	defer w.Destroy()

// 	w.SetTitle("Notas con aviso")
// 	w.SetSize(600, 600, 0)

// 	w.Bind("open", openInBrowser)
// 	w.Bind("close", func() {
// 		w.Terminate()
// 	})

// 	w.Navigate("data:text/html," + url.PathEscape(htmlContent))

// 	w.Run()
// }

// func openInBrowser(link string) error {
// 	var cmd *exec.Cmd

// 	switch runtime.GOOS {
// 	case "linux":
// 		cmd = exec.Command("xdg-open", link)
// 	case "windows":
// 		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", link)
// 	case "darwin":
// 		cmd = exec.Command("open", link)
// 	default:
// 		return fmt.Errorf("sistema operativo no soportado")
// 	}

// 	return cmd.Start()
// }


// package jobs

// import (
// 	"fmt"
// 	"net/url"
// 	"note_notifications/internal/dependencies"
// 	"note_notifications/internal/schemas"
// 	"os/exec"
// 	"runtime"
// 	"strings"
// 	"time"

// 	webview "github.com/webview/webview_go"
	
// )

// func InitNotifications(deps *dependencies.Container, noteChan chan []schemas.NoteResponse) {
// 	notes, err := deps.Services.Note.GetAllWarn()
// 	if err != nil {
// 		return
// 	}

// 	// ShowAllNotesInOneWindow(notes)

// 	for _, note := range *notes {
// 		n := note // copia local para que la goroutine no comparta el puntero
// 		go scheduleNote(n, noteChan)
// 	}
// }

// func scheduleNote(note schemas.NoteResponse, noteChan chan []schemas.NoteResponse) {
// 	// Parseamos la fecha y hora
// 	layout := "02-01-2006 15:04"
// 	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
// 	datetimeStr := fmt.Sprintf("%s %s", note.Date.String(), note.Time.String())

// 	execTime, err := time.ParseInLocation(layout, datetimeStr, loc)
// 	if err != nil {
// 		fmt.Println("Error al parsear fecha y hora:", err)
// 		return
// 	}

// 	fmt.Print(datetimeStr, " - ", execTime, " - ", time.Now(), "\n")
// 	delay := time.Until(execTime)

// 	if delay <= 0 {
// 		fmt.Println("La nota ya expiró:", note.Title)
// 		return
// 	}

// 	// Ejecutamos la nota cuando llegue el momento
// 	time.AfterFunc(delay, func() {
// 		// notes := make([]schemas.NoteResponse, 1)
// 		// notes[0] = note
// 		// showAllNotesInOneWindow(&notes)
// 		noteChan <- []schemas.NoteResponse{note}
// 	})
// }


// func ShowAllNotesInOneWindow(notes *[]schemas.NoteResponse) {
// 	if notes == nil || len(*notes) == 0 {
// 		return
// 	}

// 	var builder strings.Builder
// 	builder.WriteString(`
// 		<html>
// 		<head>
// 			<meta charset="UTF-8">
// 			<style>
// 				body {
// 					font-family: Arial, sans-serif;
// 					background-color: #f5f5f5;
// 					margin: 0;
// 					padding: 0;
// 					display: flex;
// 					flex-direction: column;
// 					height: 100vh;
// 					color: #333;
// 				}
// 				h3 {
// 					color: #007acc;
// 					text-align: center;
// 					margin: 16px 0;
// 					flex: 0 0 auto;
// 				}
// 				.grid-container {
// 					flex: 1 1 auto;
// 					overflow-y: auto;
// 					padding: 0 20px 0px 20px;
// 				}
// 				.grid {
// 					display: grid;
// 					grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
// 					gap: 15px;
// 					margin-bottom: 20px;
// 				}
// 				.note {
// 					background: #fff;
// 					border: 1px solid #ccc;
// 					border-radius: 8px;
// 					padding: 12px 16px;
// 					box-shadow: 0 2px 5px rgba(0,0,0,0.05);
// 				}
// 				.note-title {
// 					font-weight: bold;
// 					font-size: 16px;
// 					margin-bottom: 6px;
// 				}
// 				.note-meta {
// 					font-size: 13px;
// 					color: #666;
// 					margin-bottom: 6px;
// 				}
// 				.note-desc {
// 					font-size: 14px;
// 				}
// 				.actions {
// 					flex: 0 0 auto;
// 					padding: 12px 20px;
// 					background-color: #fff;
// 					border-top: 1px solid #ccc;
// 					text-align: center;
// 					position: sticky;
// 					bottom: 0;
// 					box-shadow: 0 -2px 5px rgba(0,0,0,0.05);
// 				}
// 				button {
// 					padding: 8px 16px;
// 					font-size: 14px;
// 					background-color: #007acc;
// 					color: white;
// 					border: none;
// 					border-radius: 4px;
// 					cursor: pointer;
// 				}
// 				button:hover {
// 					background-color: #005f99;
// 				}
// 				a {
// 					color: #007acc;
// 					text-decoration: none;
// 				}
// 				a:hover {
// 					text-decoration: underline;
// 				}
// 			</style>
// 			</head>
// 			<body>
// 				<h3>Notas con aviso</h3>
// 				<div class="grid-container">
// 					<div class="grid">
// 	`)

// 	for _, note := range *notes {
// 		builder.WriteString("<div class='note'>")
// 		builder.WriteString(fmt.Sprintf("<div class='note-title'>%s</div>", note.Title))
// 		builder.WriteString(fmt.Sprintf("<div class='note-meta'>📅 %s 🕒 %s</div>", note.Date.ToTime().Format("2006-01-02"), note.Time.ToTime().Format("15:04")))
// 		builder.WriteString(fmt.Sprintf("<div class='note-desc'>%s</div>", note.Description))
// 		if note.Url != nil && *note.Url != "" {
// 			builder.WriteString(fmt.Sprintf(
// 				`<br><a href="#" onclick="window.open('%s')">%s</a>`,
// 				*note.Url, *note.Url))
// 		}
// 		builder.WriteString("</div>")
// 	}

// 	builder.WriteString(`</div>`)
// 	builder.WriteString(`<div class="actions"><button onclick="window.close()">Aceptar</button></div>`)
// 	builder.WriteString(`</body></html>`)

// 	htmlContent := builder.String()

// 	debug := true
// 	w := webview.New(debug)
// 	defer w.Destroy()

// 	w.SetTitle("Notas con aviso")
// 	w.SetSize(600, 600, 0)

// 	w.Bind("open", openInBrowser)
// 	w.Bind("close", func() {
// 		w.Terminate()
// 	})

// 	w.Navigate("data:text/html," + url.PathEscape(htmlContent))

// 	w.Run()
// }

// func openInBrowser(link string) error {
// 	var cmd *exec.Cmd

// 	switch runtime.GOOS {
// 	case "linux":
// 		cmd = exec.Command("xdg-open", link)
// 	case "windows":
// 		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", link)
// 	case "darwin":
// 		cmd = exec.Command("open", link)
// 	default:
// 		return fmt.Errorf("sistema operativo no soportado")
// 	}

// 	return cmd.Start()
// }
