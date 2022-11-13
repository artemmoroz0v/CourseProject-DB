package controller

import (
	"DB/app/server"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func GetGroups(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	groups, err := server.SelectGroupsList()
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
	err = tmpl.ExecuteTemplate(w, "groups", groups)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
