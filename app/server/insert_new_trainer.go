package server

import (
	s "DB/app/model"
)

func InsertNewTrainer(newTrainer s.Trainer) error {
	_, err := db.Exec(`
		INSERT INTO Trainer 
		    (trainer_second_name, trainer_name, 
		     trainer_third_name, trainer_phone) 
		VALUES ($1, $2, $3, $4)`,
		newTrainer.TrainerSecondName,
		newTrainer.TrainerName, newTrainer.TrainerThirdName,
		newTrainer.TrainerPhone)
	return err
}
