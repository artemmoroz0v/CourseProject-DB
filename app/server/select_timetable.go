package server

import (
	s "DB/app/model"
	"log"
)

func SelectTimeTable() s.Timetable {
	newTimeTable := s.Timetable{}
	request := db.QueryRow("SELECT weekday, training_time, program_name, trainer_second_name, trainer_name, trainer_third_name FROM Timetable INNER JOIN Times USING(time_id) INNER JOIN Program USING(program_id) INNER JOIN FC_Group USING(group_id) INNER JOIN Trainer USING(trainer_id) ORDER BY weekday, training_time")
	err := request.Scan(&newTimeTable.Weekday, &newTimeTable.TrainerName, &newTimeTable.ProgramName, &newTimeTable.TrainerSecondName, &newTimeTable.TrainerName, &newTimeTable.TrainerThirdName)
	if err != nil {
		log.Fatal(err)
	}
	return newTimeTable
}
