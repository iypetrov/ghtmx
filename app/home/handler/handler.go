package handler

import (
	home "github.com/IliyaYavorovPetrov/ghtmx/app/home/server"
	"html/template"
	"log"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view/index.html"))

	homeModel := home.Server().GetUserIp(r)
	if err := tmpl.Execute(w, homeModel); err != nil {
		log.Println("error executing template :", err)
		return
	}
}
