package model

// Group symbolizes group of fitness club
type Group struct {
	GroupID       int    `json:"groupID"`
	ProgramID     int    `json:"programID"`
	Notes         string `json:"notes"`
	TrainerID     int    `json:"trainerID"`
	ClientsAmount int    `json:"clientsAmount"`
}

type Program struct {
	ProgramID   int    `json:"programID"`
	ProgramName string `json:"programName"`
}
