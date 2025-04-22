package models

import "time"

type StatusAgent struct {
	Status    string    `json:"status"`
	EditDate  time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
