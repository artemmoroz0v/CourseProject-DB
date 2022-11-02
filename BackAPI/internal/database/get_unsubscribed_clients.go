package database

import (
	s "BackAPI/internal/structs"
	"log"
)

func GetUnsubscribedClients() []s.Client {
	rows, err := db.Query("SELECT * FROM Client WHERE subscription_end < CURRENT_DATE AND subscription_end IS NULL")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	clients := []s.Client{}
	for rows.Next() {
		temp := s.Client{}
		err = rows.Scan(&temp.SubscriptionID, &temp.ClientSecondName, &temp.ClientName, &temp.ClientThirdName, &temp.Sex, &temp.Birthdate, &temp.Height, &temp.Weight, &temp.SubscriptionBegin, &temp.SubscriptionEnd)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, temp)
	}
	return clients
}
