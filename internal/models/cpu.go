package models

type Cpu struct {
	UsagePercent float64 `json:"usage_percent"`
	Temperature  float64 `json:"temperature"`
	Core         float64 `json:"core"`
	Model        float64 `json:"model"`
	Frequency    float64 `json:"frequency"`
}
