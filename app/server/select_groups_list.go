package server

import (
	s "DB/app/model"
)

func SelectGroupsList() ([]s.Group, error) {
	rows, err := db.Query(`SELECT * FROM FC_Group`)
	if err != nil {
		return nil, err
	}
	var groups []s.Group
	for rows.Next() {
		temp := s.Group{}
		err = rows.Scan(&temp.GroupID, &temp.ProgramID, &temp.Notes,
			&temp.TrainerID, &temp.ClientsAmount)
		if err != nil {
			return nil, err
		}
		groups = append(groups, temp)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return groups, nil
}
