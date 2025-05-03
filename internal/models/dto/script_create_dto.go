package dto

import "time"

type ScriptCreateDto struct {
	Command     string     `json:"command" gorm:"command"`
	Activated   bool       `json:"activated"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	ExecuteHour string     `json:"execute_hour,omitempty"`
	Weekdays    string     `json:"weekdays,omitempty"`
}
