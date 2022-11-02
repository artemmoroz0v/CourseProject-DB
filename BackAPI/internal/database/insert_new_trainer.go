package database

import (
	s "BackAPI/internal/structs"
	"log"
)

func InsertNewTrainer(newTrainer s.Trainer) {
	_, err := db.Exec("INSERT INTO Trainer (trainer_id, trainer_second_name, trainer_name, trainer_third_name, trainer_phone) VALUES ($1, $2, $3, $4, $5)",
		newTrainer.TrainerID, newTrainer.TrainerSecondName, newTrainer.TrainerName, newTrainer.TrainerThirdName, newTrainer.TrainerPhone)
	if err != nil {
		log.Fatal(err)
	}
}
