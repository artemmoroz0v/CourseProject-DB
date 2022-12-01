package model

type Client struct {
	SubscriptionID    int     `json:"subscriptionID"`
	ClientSecondName  string  `json:"clientSecondName"`
	ClientName        string  `json:"clientName"`
	ClientThirdName   string  `json:"clientThirdName"`
	Sex               string  `json:"sex"`
	Birthdate         string  `json:"birthdate"`
	Height            float64 `json:"height"`
	Weight            float64 `json:"weight"`
	SubscriptionBegin string  `json:"subscriptionBegin"`
	SubscriptionEnd   string  `json:"subscriptionEnd"`
}
