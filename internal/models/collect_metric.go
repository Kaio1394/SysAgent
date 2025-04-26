package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CollectMetric struct {
	Uuid       string    `json:"uuid" gorm:"primaryKey"`
	JsonResult string    `json:"json_result" gorm:"json_result"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (m *CollectMetric) BeforeCreate(tx *gorm.DB) (err error) {
	m.Uuid = uuid.New().String()
	return
}
