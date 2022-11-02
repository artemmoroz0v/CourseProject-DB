package database

import (
	s "BackAPI/internal/structs"
	"log"
)

func InsertNewClient(newClient s.Client) {
	_, err := db.Exec("INSERT INTO Client (subscription_id, client_second_name, client_name, client_third_name, sex, birthdate, height, weight, subscription_begin, subscription_end) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		newClient.SubscriptionID, newClient.ClientSecondName, newClient.ClientName, newClient.ClientThirdName, newClient.Sex, newClient.Birthdate, newClient.Height, newClient.Weight, newClient.SubscriptionBegin, newClient.SubscriptionEnd)
	if err != nil {
		log.Fatal(err)
	}
}
