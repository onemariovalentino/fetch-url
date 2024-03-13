package models

import "time"

type (
	Metadata struct {
		URL       string    `json:"url"`
		NumLinks  int       `json:"num_links"`
		NumImages int       `json:"num_images"`
		LastFetch time.Time `json:"last_fetch"`
	}
)
