package database

import (
	s "BackAPI/internal/structs"
	"log"
)

func GetTrainersList() []s.Trainer {
	rows, err := db.Query("SELECT * FROM Trainer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	trainers := []s.Trainer{}
	for rows.Next() {
		temp := s.Trainer{}
		err = rows.Scan(&temp.TrainerID, &temp.TrainerSecondName, &temp.TrainerName, &temp.TrainerThirdName, &temp.TrainerPhone)
		if err != nil {
			log.Fatal(err)
		}
		trainers = append(trainers, temp)
	}
	return trainers
}
