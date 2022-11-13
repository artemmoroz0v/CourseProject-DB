package controller

import (
	"DB/app/model"
	"DB/app/server"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func PostTrainer(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newTrainer model.Trainer
	newTrainer.TrainerSecondName = r.FormValue("trainer_second_name")
	newTrainer.TrainerName = r.FormValue("trainer_name")
	newTrainer.TrainerThirdName = r.FormValue("trainer_third_name")
	newTrainer.TrainerPhone = r.FormValue("trainer_phone")
	err := server.InsertNewTrainer(newTrainer)
	if err != nil {
		log.Fatal(err)
	}
}
