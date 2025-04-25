package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Agent struct {
	Uuid      string    `json:"uuid" gorm:"primary_key"`
	Status    string    `json:"status" gorm:"status"`
	EditDate  time.Time `gorm:"autoUpdateTime"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (m *Agent) BeforeCreate(tx *gorm.DB) (err error) {
	m.Uuid = uuid.New().String()
	return
}
