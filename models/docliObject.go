package models

type DocliObject struct {
	FullName string `json:"full_name"`
	UserId string `json:"user_id"`
	OriginalName string `json:"image_name"`
	UniqueId string `json:"unique_id"`
	Ports []PortObject `json:"ports"`
	Networks []string `json:"networks"`
	Volumes []string `json:"volumes"`
}

type PortObject struct {
	Container int `json:"container"`
	Host int `json:"host"`
}
