package controller

import (
	s "DB/app/model"
	"DB/app/server"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func SelectTimeTableForEachGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	groupID, _ := strconv.Atoi(r.FormValue("select_group_id"))
	timetables, err := server.TimeTableEachGroup(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "timetable.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "timetable", timetables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SelectTimeTableForEachProgram(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	programID, _ := strconv.Atoi(r.FormValue("select_program_id"))
	timetables, err := server.TimeTableEachProgram(programID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "timetable.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "timetable", timetables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SelectTimeTableForEachTrainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	trainerID, _ := strconv.Atoi(r.FormValue("select_trainer_id"))
	timetables, err := server.TimeTableEachTrainer(trainerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := filepath.Join("public", "pages", "timetable.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = tmpl.ExecuteTemplate(w, "timetable", timetables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func InsertProgramInTimeTable(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	newTimeTable := s.Timetable{}
	newTimeTable.GroupID, _ = strconv.Atoi(r.FormValue("insert_group_id"))
	newTimeTable.Weekday = r.FormValue("insert_weekday")
	newTimeTable.TrainingTime = r.FormValue("insert_training_time")
	err := server.InsertInTimeTable(newTimeTable)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteProgramFromTimeTable(w http.ResponseWriter, r *http.Request,
	p httprouter.Params) {
	id, err := strconv.Atoi(r.FormValue("delete_group_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = server.DeleteFromTimeTable(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
