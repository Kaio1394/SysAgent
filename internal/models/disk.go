package models

type Disk struct {
	TotalSpace float64 `json:"total_space"`
	UsedSpace  float64 `json:"used_space"`
	FreeSpace  float64 `json:"free_space"`
}
