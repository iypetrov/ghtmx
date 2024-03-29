package ip

import (
	"html/template"
	"log"
	"net/http"
)

func GetRequestIPHandler(server *Server) func(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		requestIPModel := server.GetUserIp(r)
		requestIPResponseDTO := RequestIPModelToRequestIPResponseDTO(requestIPModel)
		if err := tmpl.Execute(w, requestIPResponseDTO); err != nil {
			log.Println("error executing template :", err)
			return
		}
	}
}
