package models

type ExecutionHistory struct {
}

func (e *ExecutionHistory) TableName() string {
	return "th_execution_history"
}
