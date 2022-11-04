package database

import (
	s "BackAPI/internal/structs"
)

func SelectGroupList(number int) ([]s.Client, error) {
	rows, err := db.Query("SELECT subscription_id, client_second_name, client_name, client_third_name, sex, birthdate, height, weight, subscription_begin, subscription_end FROM Client INNER JOIN group_client USING(subscription_id) WHERE group_id = $1", number)
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
