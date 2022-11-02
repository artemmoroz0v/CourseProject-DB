package database

import (
	s "BackAPI/internal/structs"
	"log"
)

func GetGroupList(number int) []s.Client {
	rows, err := db.Query("SELECT subscription_id, client_second_name, client_name, client_third_name, sex, birthdate, height, weight, subscription_begin, subscription_end FROM Client INNER JOIN Client USING(subscription_id) INNER JOIN FC_Group USING(group_id) WHERE group_id = $1))", number)
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
