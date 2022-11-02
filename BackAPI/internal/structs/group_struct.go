package structs

type Group struct {
	GroupID       int    `json:"groupID"`
	ProgramID     int    `json:"programID"`
	Notes         string `json:"notes"`
	TrainerID     int    `json:"trainerID"`
	ClientsAmount int    `json:"clientsAmount"`
}
