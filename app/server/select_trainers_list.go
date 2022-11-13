package server

import (
	s "DB/app/model"
)

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
