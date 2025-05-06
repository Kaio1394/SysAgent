package models

type ExecutionHistory struct {
	Uuid         string `json:"uuid" gorm:"primaryKey"`
	UuidExecuted string `json:"uuid_executed"`
	Date         string `json:"date"`
	Result       string `json:"result"`
}

func (e *ExecutionHistory) TableName() string {
	return "th_execution_history"
}
