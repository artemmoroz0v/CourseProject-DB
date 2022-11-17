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

func SelectGroups(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Groups, err := server.SelectGroupsList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "groups.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		Groups []model.Group
		Group  []model.Client
	}{
		Groups,
		nil,
	}
	err = tmpl.ExecuteTemplate(w, "data", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SelectGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	groupID, err := strconv.Atoi(r.FormValue("group_choose"))
	if err != nil {
		return
	}
	Group, err := server.SelectGroupList(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Groups, err := server.SelectGroupsList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "groups.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		Groups []model.Group
		Group  []model.Client
	}{
		Groups,
		Group,
	}
	err = tmpl.ExecuteTemplate(w, "data", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertClientIntoGroup(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	clientID, err := strconv.Atoi(r.FormValue("client_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	groupID, err := strconv.Atoi(r.FormValue("group_append"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.InsertClientIntoGroup(clientID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
