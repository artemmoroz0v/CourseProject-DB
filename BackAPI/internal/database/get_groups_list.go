package database

import (
	s "BackAPI/internal/structs"
	"database/sql"
	"log"
)

func GetGroupsList() []s.Group {
	rows, err := db.Query("SELECT * FROM FC_Group")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {

		}
	}(rows)
	var groups []s.Group
	for rows.Next() {
		temp := s.Group{}
		err = rows.Scan(&temp.GroupID, &temp.ProgramID, &temp.TrainerID)
		if err != nil {
			log.Fatal(err)
		}
		groups = append(groups, temp)
	}
	return groups
}
