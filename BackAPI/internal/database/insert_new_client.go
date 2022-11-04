package database

import (
	s "BackAPI/internal/structs"
)

func InsertNewClient(newClient s.Client) error {
	_, err := db.Exec("INSERT INTO Client (client_second_name, client_name, client_third_name, sex, birthdate, height, weight, subscription_begin, subscription_end) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		newClient.ClientSecondName, newClient.ClientName, newClient.ClientThirdName, newClient.Sex, newClient.Birthdate, newClient.Height, newClient.Weight, newClient.SubscriptionBegin, newClient.SubscriptionEnd)
	return err
}
