package models

type Metric struct {
	Memory Memory `json:"memory"`
	Cpu    Cpu    `json:"cpu"`
	Disk   Disk   `json:"disk"`
}
