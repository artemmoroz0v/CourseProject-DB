package API

import (
	"BackAPI/internal/database"
	s "BackAPI/internal/structs"
	"log"
	"net/http"
)

func PostTrainer(w http.ResponseWriter, r *http.Request) {
	var newTrainer s.Trainer
	newTrainer.TrainerSecondName = r.FormValue("trainer_second_name")
	newTrainer.TrainerName = r.FormValue("trainer_name")
	newTrainer.TrainerThirdName = r.FormValue("trainer_third_name")
	newTrainer.TrainerPhone = r.FormValue("trainer_phone")
	err := database.InsertNewTrainer(newTrainer)
	if err != nil {
		log.Fatal(err)
	}
}
