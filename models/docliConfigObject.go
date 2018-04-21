package models

type DocliConfigObject struct {
	FullName string `json:"full_name"`
	UserId string `json:"user_id"`
	OriginalName string `json:"image_name"`
	UniqueId string `json:"unique_id"`
	Ports []int `json:"ports"`
	Networks []string `json:"networks"`
	Volumes []string `json:"volumes"`
}
