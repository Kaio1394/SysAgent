package models

type Cpu struct {
	Model        string  `json:"model"`
	UsagePercent float64 `json:"usage_percent"`
	Core         int32   `json:"core"`
	Frequency    float64 `json:"frequency"`
}
