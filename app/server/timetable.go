package server

import (
	s "DB/app/model"
)

func TimeTableEachGroup(groupID int) ([]s.Timetable, error) {
	rows, err := db.Query(`
	SELECT group_id, weekday, training_time 
	FROM timetable 
	INNER JOIN times 
	USING(time_id) 
	WHERE group_id = $1`,
		groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timetables := []s.Timetable{}
	for rows.Next() {
		temp := s.Timetable{}
		err = rows.Scan(&temp.GroupID, &temp.Weekday, &temp.TrainingTime)
		if err != nil {
			return nil, err
		}
		timetables = append(timetables, temp)
	}
	return timetables, nil
}

func TimeTableEachProgram(programID int) ([]s.Timetable, error) {
	rows, err := db.Query(`
	SELECT group_id, weekday, training_time 
	FROM timetable 
	INNER JOIN times 
	USING(time_id) 
	WHERE program_id = $1`,
		programID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timetables := []s.Timetable{}
	for rows.Next() {
		temp := s.Timetable{}
		err = rows.Scan(&temp.GroupID, &temp.Weekday, &temp.TrainingTime)
		if err != nil {
			return nil, err
		}
		timetables = append(timetables, temp)
	}
	return timetables, nil
}

func TimeTableEachTrainer(trainerID int) ([]s.Timetable, error) {
	rows, err := db.Query(`
	SELECT group_id, weekday, training_time 
	FROM timetable 
	INNER JOIN times 
	USING(time_id) 
	INNER JOIN fc_group
	USING(group_id)
	WHERE trainer_id = $1`,
		trainerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	timetables := []s.Timetable{}
	for rows.Next() {
		temp := s.Timetable{}
		err = rows.Scan(&temp.GroupID, &temp.Weekday, &temp.TrainingTime)
		if err != nil {
			return nil, err
		}
		timetables = append(timetables, temp)
	}
	return timetables, nil
}

func InsertInTimeTable(newTimeTable s.Timetable) error {
	_, err := db.Exec(`
	INSERT INTO timetable (group_id, weekday, training_time) 
	VALUES ($1, $2, $3)`,
		newTimeTable.GroupID, newTimeTable.Weekday, newTimeTable.TrainingTime)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromTimeTable(groupid int) error {
	_, err := db.Exec(`
	DELETE FROM timetable 
	WHERE group_id = $1`,
		groupid)
	if err != nil {
		return err
	}
	return nil
}
