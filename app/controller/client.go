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
	if err = tmpl.ExecuteTemplate(w, "data", data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newClient model.Client
	newClient.ClientSecondName = r.FormValue("insert_second_name")
	newClient.ClientName = r.FormValue("insert_name")
	newClient.ClientThirdName = r.FormValue("insert_third_name")
	newClient.Sex = r.FormValue("insert_sex")
	newClient.Birthdate = r.FormValue("insert_birthdate")
	newClient.Height, _ = strconv.ParseFloat(r.FormValue("insert_height"), 64)
	newClient.Weight, _ = strconv.ParseFloat(r.FormValue("insert_weight"), 64)
	newClient.SubscriptionBegin = r.FormValue("insert_subscription_begin")
	newClient.SubscriptionEnd = r.FormValue("insert_subscription_end")
	if err := server.InsertNewClient(newClient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdateClientSubscription(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("update_subscription_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	date := r.FormValue("update_subscription_subscription_end")
	if err = server.UpdateClientSubscription(id, date); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdateHeightAndWeight(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("update_height&weight_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newHeight, err :=
		strconv.ParseFloat(r.FormValue("update_height&weight_height"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newWeight, err :=
		strconv.ParseFloat(r.FormValue("update_height&weight_weight"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = server.UpdateHeightAndWeight(id, newHeight, newWeight); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteClient(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("delete_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = server.DeleteClient(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
