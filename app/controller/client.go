package controller

import (
	"DB/app/model"
	"DB/app/server"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

func SelectClients(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
	path := filepath.Join("public", "pages", "client.html")
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

func InsertClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newClient model.Client
	newClient.ClientSecondName = r.FormValue("client_second_name")
	newClient.ClientName = r.FormValue("client_name")
	newClient.ClientThirdName = r.FormValue("client_third_name")
	newClient.Sex = r.FormValue("sex")
	newClient.Birthdate = r.FormValue("birthdate")
	newClient.Height, _ = strconv.ParseFloat(r.FormValue("height"), 64)
	newClient.Weight, _ = strconv.ParseFloat(r.FormValue("weight"), 64)
	newClient.SubscriptionBegin = r.FormValue("subscription_begin")
	newClient.SubscriptionEnd = r.FormValue("subscription_end")
	err := server.InsertNewClient(newClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdateClientSubscription(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("client_id_update"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	date := r.FormValue("subscription_end_update")
	err = server.UpdateClientSubscription(id, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdateHeightAndWeight(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("client_id_update_hw"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newHeight, err :=
		strconv.ParseFloat(r.FormValue("client_params_update_height"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newWeight, err :=
		strconv.ParseFloat(r.FormValue("client_params_update_weight"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.UpdateHeightAndWeight(id, newHeight, newWeight)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteClient(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("client_id_delete"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.DeleteClient(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
