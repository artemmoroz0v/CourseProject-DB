package database

import (
	s "BackAPI/internal/structs"
)

func SelectUnsubscribedClients() ([]s.Client, error) {
	rows, err := db.Query("SELECT * FROM Client WHERE subscription_end < CURRENT_DATE AND subscription_end IS NULL")
	if err != nil {
		return nil, err
	}
	var clients []s.Client
	for rows.Next() {
		temp := s.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName, &temp.ClientName, &temp.ClientThirdName, &temp.Sex, &temp.Birthdate, &temp.Height, &temp.Weight, &temp.SubscriptionBegin, &temp.SubscriptionEnd)
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
