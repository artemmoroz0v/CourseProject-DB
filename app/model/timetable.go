package model

type Timetable struct {
	GroupID      int    `json:"groupID"`
	Weekday      string `json:"weekday"`
	TrainingTime string `json:"trainingTime"`
}
