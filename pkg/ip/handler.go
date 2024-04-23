package ip

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func CreateRequestIPHandler(s *Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		model, err := s.CreateRequestIPModel(r)
		if err != nil {
			http.Error(w, "Couldn't save the IP of the user", http.StatusInternalServerError)
			return
		}

		dto := RequestIPModelToRequestIPResponseDTO(model)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dto); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}

func GetRequestIPHandler(s *Server, dbRunning bool) func(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.RemoteAddr, ":")
		dto := RequestIPResponseDTO{
			IP: parts[0],
		}

		data := struct {
			DTO       RequestIPResponseDTO
			DBRunning bool
		}{
			DTO:       dto,
			DBRunning: dbRunning,
		}
		if err := tmpl.Execute(w, data); err != nil {
			log.Println("error executing template :", err)
			return
		}
	}
}

func GetStatsIPHandler(s *Server) func(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/stats.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		models, err := s.GetStatsIPModels()
		if err != nil {
			return
		}

		dtos := StatsIPModelsToStatsIPResponseDTOs(models)
		if err := tmpl.Execute(w, dtos); err != nil {
			log.Println("error executing template :", err)
			return
		}
	}
}
