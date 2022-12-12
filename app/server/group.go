package server

import "DB/app/model"

func InsertNewGroup(newGroup model.Group) error {
	_, err := db.Exec(`
		INSERT INTO FC_Group 
		    (group_id, program_id, notes, trainer_id, clients_amount) 
		VALUES ($1, $2, $3, $4, $5)`,
		newGroup.GroupID, newGroup.ProgramID, newGroup.Notes,
		newGroup.TrainerID, newGroup.ClientsAmount)
	return err
}

//inserting new group in database

func SelectGroupList(number int) ([]model.Client, error) {
	rows, err := db.Query(`
		SELECT subscription_id, client_second_name, client_name, 
		       client_third_name, sex, birthdate, height, weight, 
		       subscription_begin, subscription_end 
		FROM Client 
		    INNER JOIN group_client USING(subscription_id) 
		WHERE group_id = $1
		ORDER BY subscription_id DESC`, number)
	if err != nil {
		return nil, err
	}
	var clients []model.Client
	for rows.Next() {
		temp := model.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName,
			&temp.ClientName, &temp.ClientThirdName, &temp.Sex, &temp.Birthdate,
			&temp.Height, &temp.Weight, &temp.SubscriptionBegin,
			&temp.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
		temp.Birthdate = temp.Birthdate[:10]
		temp.SubscriptionBegin = temp.SubscriptionBegin[:10]
		temp.SubscriptionEnd = temp.SubscriptionEnd[:10]
		clients = append(clients, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return clients, nil
}

//selecting group by it's id from database

func SelectGroupsList() ([]model.Group, error) {
	rows, err := db.Query(`
		SELECT * FROM FC_Group
		ORDER BY group_id DESC`)
	if err != nil {
		return nil, err
	}
	var groups []model.Group
	for rows.Next() {
		temp := model.Group{}
		err = rows.Scan(&temp.GroupID, &temp.ProgramID, &temp.Notes,
			&temp.TrainerID, &temp.ClientsAmount)
		if err != nil {
			return nil, err
		}
		groups = append(groups, temp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return groups, nil
}

//selecting all groups from database

func InsertClientIntoGroup(clientID int, groupID int) error {
	_, err := db.Exec(`
		INSERT INTO group_client (group_id, subscription_id)
		VALUES ($1, $2)`,
		groupID, clientID)
	return err
}

//inserting new client in group by ids

func DeleteClientFromGroup(clientID int, groupID int) error {
	_, err := db.Exec(`
		DELETE FROM group_client
		WHERE group_id = $1 AND subscription_id = $2`,
		groupID, clientID)
	return err
}

//deleting client from group database

func DeleteGroup(id int) error {
	_, err := db.Exec(`
		DELETE FROM FC_Group
		WHERE (group_id = $1)`,
		id)
	return err
}

//deleting group by it's id from database
