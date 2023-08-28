package calendar

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/flpnascto/lecture-calendar/models"
)

func GetEvents() []models.Calendar {
	events, err := readJsonCalendar()
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	return events
}

func readJsonCalendar() ([]models.Calendar, error) {
	file := filepath.Join(".", "calendar", "calendar.json")
	var events []models.Calendar

	jsonData, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo - ", err)
		return events, err
	}

	err = json.Unmarshal(jsonData, &events)
	if err != nil {
		fmt.Println("Erro ao decodificar o JSON - ", err)
		return events, err
	}

	return events, nil
}
