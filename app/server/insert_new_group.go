package server

import (
	s "DB/app/model"
)

func InsertNewGroup(newGroup s.Group) error {
	_, err := db.Exec(`
		INSERT INTO FC_Group 
		    (group_id, program_id, notes, trainer_id, clients_amount) 
		VALUES ($1, $2, $3, $4, $5)`,
		newGroup.GroupID, newGroup.ProgramID, newGroup.Notes,
		newGroup.TrainerID, newGroup.ClientsAmount)
	return err
}
