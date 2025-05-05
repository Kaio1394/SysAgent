package dto

type ScriptCreateDto struct {
	Command     string `json:"command" gorm:"command"`
	Activated   bool   `json:"activated"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	ExecuteHour string `json:"execute_hour,omitempty"`
	Weekdays    string `json:"weekdays,omitempty"`
}
