package model

type Trainer struct {
	TrainerID         int    `json:"trainerID"`
	TrainerSecondName string `json:"trainerSecondName"`
	TrainerName       string `json:"trainerName"`
	TrainerThirdName  string `json:"trainerThirdName"`
	TrainerPhone      string `json:"trainerPhone"`
}
