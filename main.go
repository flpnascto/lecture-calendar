package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/flpnascto/lecture-calendar/calendar"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join(".", "public", "calendar.html")
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
		return
	}
	events := calendar.GetEvents()

	t.Execute(w, events)

}
