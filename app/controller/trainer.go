package controller

import (
	"DB/app/model"
	"DB/app/server"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func SelectTrainers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	trainers, err := server.SelectTrainersList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "trainer.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "trainers", trainers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertTrainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newTrainer model.Trainer
	newTrainer.TrainerSecondName = r.FormValue("trainer_second_name")
	newTrainer.TrainerName = r.FormValue("trainer_name")
	newTrainer.TrainerThirdName = r.FormValue("trainer_third_name")
	newTrainer.TrainerPhone = r.FormValue("trainer_phone")
	err := server.InsertNewTrainer(newTrainer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
