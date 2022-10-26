package structs

import "time"

type Client struct {
	SubscriptionID    int       `json:"subscriptionID"`
	ClientSecondName  string    `json:"clientSecondName"`
	ClientName        string    `json:"clientName"`
	ClientThirdName   string    `json:"clientThirdName"`
	Sex               string    `json:"sex"`
	Birthdate         time.Time `json:"birthdate"`
	Height            float64   `json:"height"`
	Weight            float64   `json:"weight"`
	SubscriptionBegin time.Time `json:"subscriptionBegin"`
	SubscriptionEnd   time.Time `json:"subscriptionEnd"`
}
