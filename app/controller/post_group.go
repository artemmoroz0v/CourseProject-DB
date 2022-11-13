package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func PostGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	/*var newGroup model.Group
	var err error
	newGroup.ProgramID, err = strconv.Atoi(r.FormValue("group_program_list"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newGroup.Notes = r.FormValue("group_note")
	newGroup.TrainerID, err = strconv.Atoi(r.FormValue("group_trainer_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newGroup.ClientsAmount, err =
		strconv.Atoi(r.FormValue("group_clients_amount"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.InsertNewGroup(newGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}*/
}
