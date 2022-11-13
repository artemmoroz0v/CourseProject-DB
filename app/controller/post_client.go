package controller

import (
	"DB/app/model"
	"DB/app/server"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func PostClient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newClient model.Client
	newClient.ClientSecondName = r.FormValue("client_second_name")
	newClient.ClientName = r.FormValue("client_name")
	newClient.ClientThirdName = r.FormValue("client_third_name")
	newClient.Sex = r.FormValue("sex")
	//newClient.Birthdate = r.FormValue("birthdate")
	//newClient.Height = r.FormValue("height")
	//newClient.Weight = r.FormValue("weight")
	//newClient.SubscriptionBegin = r.FormValue("subscription_begin")
	//newClient.SubscriptionEnd = r.FormValue("subscription_end")
	err := server.InsertNewClient(newClient)
	if err != nil {
		log.Fatal(err)
	}
}
