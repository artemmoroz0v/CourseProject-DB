package controller

import (
	"DB/app/model"
	"DB/app/server"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var timetableByGroup = 0
var timetableByProgram = 0
var timetableByTrainer = 0

func SelectTimetable(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	byGroup, err := server.SelectTimetableByGroup(timetableByGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	byProgram, err := server.SelectTimetableByProgram(timetableByProgram)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	byTrainer, err := server.SelectTimetableByTrainer(timetableByTrainer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := struct {
		ByGroup   []model.Timetable
		ByProgram []model.Timetable
		ByTrainer []model.Timetable
	}{
		byGroup,
		byProgram,
		byTrainer,
	}
	path := filepath.Join("public", "pages", "timetable.html")
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

func SelectTimetableByGroup(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	tempStr := r.FormValue("select_group_id")
	if tempStr != "" {
		timetableByGroup, _ = strconv.Atoi(tempStr)
	}
	printAnswer(w, successRes, successAns)
}

func SelectTimetableByProgram(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	tempStr := r.FormValue("select_program_id")
	if tempStr != "" {
		timetableByProgram, _ = strconv.Atoi(tempStr)
	}
	printAnswer(w, successRes, successAns)
}

func SelectTimetableByTrainer(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	tempStr := r.FormValue("select_trainer_id")
	if tempStr != "" {
		timetableByTrainer, _ = strconv.Atoi(tempStr)
	}
	printAnswer(w, successRes, successAns)
}

func InsertTimetable(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	var newTimetable model.Timetable
	newTimetable.GroupID, _ = strconv.Atoi(r.FormValue("insert_group_id"))
	newTimetable.Weekday = r.FormValue("insert_weekday")
	newTimetable.TrainingTime = r.FormValue("insert_training_time")
	if err := server.InsertTimetable(newTimetable); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}

func DeleteTimetable(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	var newTimetable model.Timetable
	newTimetable.GroupID, _ = strconv.Atoi(r.FormValue("insert_group_id"))
	newTimetable.Weekday = r.FormValue("insert_weekday")
	newTimetable.TrainingTime = r.FormValue("insert_training_time")
	if err := server.DeleteTimetable(newTimetable); err != nil {
		printAnswer(w, errorRes, err.Error())
		return
	}
	printAnswer(w, successRes, successAns)
}
