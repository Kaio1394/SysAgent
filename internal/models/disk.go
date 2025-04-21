package models

type Disk struct {
	TotalSpace string `json:"total_space"`
	UsedSpace  string `json:"used_space"`
	FreeSpace  string `json:"free_space"`
}
