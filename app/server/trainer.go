package server

import s "DB/app/model"

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

func SelectTrainersList() ([]s.Trainer, error) {
	rows, err := db.Query(`SELECT * FROM Trainer`)
	if err != nil {
		return nil, err
	}
	var trainers []s.Trainer
	for rows.Next() {
		temp := s.Trainer{}
		err = rows.Scan(&temp.TrainerID, &temp.TrainerSecondName,
			&temp.TrainerName, &temp.TrainerThirdName, &temp.TrainerPhone)
		if err != nil {
			return nil, err
		}
		trainers = append(trainers, temp)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return trainers, nil
}
