package ip

import (
	"html/template"
	"log"
	"net/http"
)

func GetRequestIPHandler(server *Server) func(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		requestIPModel, err := server.GetUserIp(r)
		if err != nil {
			http.Error(w, "Failed to get user IP: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, requestIPModel); err != nil {
			log.Println("error executing template :", err)
			return
		}
	}
}
