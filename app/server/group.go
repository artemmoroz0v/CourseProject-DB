package server

import s "DB/app/model"

func InsertNewGroup(newGroup s.Group) error {
	_, err := db.Exec(`
		INSERT INTO FC_Group 
		    (group_id, program_id, notes, trainer_id, clients_amount) 
		VALUES ($1, $2, $3, $4, $5)`,
		newGroup.GroupID, newGroup.ProgramID, newGroup.Notes,
		newGroup.TrainerID, newGroup.ClientsAmount)
	return err
}

func SelectGroupList(number int) ([]s.Client, error) {
	rows, err := db.Query(`
		SELECT subscription_id, client_second_name, client_name, 
		       client_third_name, sex, birthdate, height, weight, 
		       subscription_begin, subscription_end 
		FROM Client 
		    INNER JOIN group_client USING(subscription_id) 
		WHERE group_id = $1`, number)
	if err != nil {
		return nil, err
	}
	var clients []s.Client
	for rows.Next() {
		temp := s.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName,
			&temp.ClientName, &temp.ClientThirdName, &temp.Sex, &temp.Birthdate,
			&temp.Height, &temp.Weight, &temp.SubscriptionBegin,
			&temp.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
		clients = append(clients, temp)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return clients, nil
}

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
