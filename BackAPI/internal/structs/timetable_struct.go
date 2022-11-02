package structs

type Timetable struct {
	Weekday           int    `json:"weekday"`
	TrainingTime      string `json:"trainingTime"`
	ProgramName       string `json:"programName"`
	TrainerSecondName string `json:"trainerSecondName"`
	TrainerName       string `json:"trainerName"`
	TrainerThirdName  string `json:"trainerThirdName"`
}
