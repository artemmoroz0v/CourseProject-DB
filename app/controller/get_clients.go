package controller

import (
	"DB/app/model"
	"DB/app/server"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func GetClients(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	clients, err := server.SelectClients()
	unsubscribedClients, err := server.SelectUnsubscribedClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		Clients             []model.Client
		UnsubscribedClients []model.Client
	}{
		clients,
		unsubscribedClients,
	}
	path := filepath.Join("public", "pages", "clients.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "data", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
