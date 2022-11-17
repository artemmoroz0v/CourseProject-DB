package server

import "DB/app/model"

func InsertNewClient(newClient model.Client) error {
	_, err := db.Exec(`
		INSERT INTO Client 
		    (client_second_name, client_name, client_third_name, sex, 
		     birthdate, height, weight, subscription_begin, subscription_end) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		newClient.ClientSecondName, newClient.ClientName,
		newClient.ClientThirdName, newClient.Sex, newClient.Birthdate,
		newClient.Height, newClient.Weight,
		newClient.SubscriptionBegin, newClient.SubscriptionEnd)
	return err
}

func SelectClients() ([]model.Client, error) {
	rows, err := db.Query(`SELECT * FROM Client`)
	if err != nil {
		return nil, err
	}
	var clients []model.Client
	for rows.Next() {
		temp := model.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName,
			&temp.ClientName, &temp.ClientThirdName, &temp.Sex,
			&temp.Birthdate, &temp.Height, &temp.Weight,
			&temp.SubscriptionBegin, &temp.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
		temp.Birthdate = temp.Birthdate[:10]
		temp.SubscriptionBegin = temp.SubscriptionBegin[:10]
		temp.SubscriptionEnd = temp.SubscriptionEnd[:10]
		clients = append(clients, temp)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func SelectUnsubscribedClients() ([]model.Client, error) {
	rows, err := db.Query(`
		SELECT * 
		FROM Client 
		WHERE subscription_end < CURRENT_DATE`)
	if err != nil {
		return nil, err
	}
	var clients []model.Client
	for rows.Next() {
		temp := model.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName,
			&temp.ClientName, &temp.ClientThirdName, &temp.Sex,
			&temp.Birthdate, &temp.Height, &temp.Weight,
			&temp.SubscriptionBegin, &temp.SubscriptionEnd)
		if err != nil {
			return nil, err
		}
		temp.Birthdate = temp.Birthdate[:10]
		temp.SubscriptionBegin = temp.SubscriptionBegin[:10]
		temp.SubscriptionEnd = temp.SubscriptionEnd[:10]
		clients = append(clients, temp)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func UpdateSubscription(id int, date string) error {
	_, err := db.Exec(`
		UPDATE client
		SET subscription_end = $1
		WHERE subscription_id = $2`,
		date, id)
	return err
}
