package database

import (
	s "BackAPI/internal/structs"
)

func SelectGroupsList() ([]s.Group, error) {
	rows, err := db.Query("SELECT * FROM FC_Group")
	if err != nil {
		return nil, err
	}
	var groups []s.Group
	for rows.Next() {
		temp := s.Group{}
		err = rows.Scan(&temp.GroupID, &temp.ProgramID, &temp.TrainerID)
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
