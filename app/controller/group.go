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

var selectedGroup = 0

func SelectGroups(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	groups, err := server.SelectGroupsList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	group, err := server.SelectGroupList(selectedGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		Groups []model.Group
		Group  []model.Client
	}{
		groups,
		group,
	}
	path := filepath.Join("public", "pages", "group.html")
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

func SelectGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tempStr := r.FormValue("select_group_id")
	if tempStr != "" {
		selectedGroup, _ = strconv.Atoi(tempStr)
	}
	printAnswer(w, successRes, successAns)
}

func InsertGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newGroup model.Group
	var err error
	newGroup.ProgramID, err = strconv.Atoi(r.FormValue("insert_program_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newGroup.Notes = r.FormValue("insert_notes")
	newGroup.TrainerID, _ = strconv.Atoi(r.FormValue("insert_trainer_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	newGroup.ClientsAmount, err =
		strconv.Atoi(r.FormValue("insert_clients_amount"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.InsertNewGroup(newGroup); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

func InsertClientIntoGroup(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	clientID, err := strconv.Atoi(r.FormValue("client_insert_client_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	groupID, err := strconv.Atoi(r.FormValue("client_insert_group_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.InsertClientIntoGroup(clientID, groupID); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

func DeleteClientFromGroup(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	clientID, err := strconv.Atoi(r.FormValue("client_delete_client_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	groupID, err := strconv.Atoi(r.FormValue("client_delete_group_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.DeleteClientFromGroup(clientID, groupID); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("delete_group_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.DeleteGroup(id); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
