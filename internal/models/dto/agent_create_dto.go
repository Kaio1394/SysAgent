package dto

type AgentCreateDto struct {
	Status string `json:"status" gorm:"primaryKey"`
}
