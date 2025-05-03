package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Script struct {
	Uuid      string `json:"uuid" gorm:"primaryKey"`
	Command   string `json:"command" gorm:"command"`
	Activated bool   `json:"activated"`

	// fields to execution script
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	ExecuteHour string     `json:"execute_hour,omitempty"`
	Weekdays    string     `json:"weekdays,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	EditedAt  time.Time `gorm:"autoUpdateTime"`
}

func (s *Script) TableName() string {
	return "tb_scripts"
}

func (s *Script) BeforeCreate(tx *gorm.DB) (err error) {
	s.Uuid = uuid.New().String()
	return
}
