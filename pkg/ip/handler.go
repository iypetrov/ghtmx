package ip

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
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

func GetRequestIPHandler(s *Server) func(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		dto := RequestIPResponseDTO{
			IP: r.RemoteAddr,
		}
		if err := tmpl.Execute(w, dto); err != nil {
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
