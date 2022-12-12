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
	if err = tmpl.ExecuteTemplate(w, "trainers", trainers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertTrainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newTrainer model.Trainer
	newTrainer.TrainerSecondName = r.FormValue("insert_second_name")
	newTrainer.TrainerName = r.FormValue("insert_name")
	newTrainer.TrainerThirdName = r.FormValue("insert_third_name")
	newTrainer.TrainerPhone = r.FormValue("insert_phone")
	if err := server.InsertNewTrainer(newTrainer); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

func DeleteTrainer(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("delete_id"))
	if err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	if err = server.DeleteTrainer(id); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
