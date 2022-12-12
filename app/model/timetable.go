package model

// Timetable symbolizes a part of timetable of fitness club
type Timetable struct {
	GroupID      int    `json:"groupID"`
	Weekday      string `json:"weekday"`
	TrainingTime string `json:"trainingTime"`
}
