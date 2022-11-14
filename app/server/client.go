package server

import s "DB/app/model"

func InsertNewClient(newClient s.Client) error {
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

func SelectClients() ([]s.Client, error) {
	rows, err := db.Query(`SELECT * FROM Client`)
	if err != nil {
		return nil, err
	}
	var clients []s.Client
	for rows.Next() {
		temp := s.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName,
			&temp.ClientName, &temp.ClientThirdName, &temp.Sex,
			&temp.Birthdate, &temp.Height, &temp.Weight,
			&temp.SubscriptionBegin, &temp.SubscriptionEnd)
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

func SelectUnsubscribedClients() ([]s.Client, error) {
	rows, err := db.Query(`
		SELECT * 
		FROM Client 
		WHERE subscription_end < CURRENT_DATE`)
	if err != nil {
		return nil, err
	}
	var clients []s.Client
	for rows.Next() {
		temp := s.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName,
			&temp.ClientName, &temp.ClientThirdName, &temp.Sex,
			&temp.Birthdate, &temp.Height, &temp.Weight,
			&temp.SubscriptionBegin, &temp.SubscriptionEnd)
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
