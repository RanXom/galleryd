package api

import "time"

type photoResponse struct {
	ID           string    `json:"id"`
	DateTaken    time.Time `json:"dateTaken"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	ThumbnailURL string    `json:"thumbnail"`
	PhotoURL     string    `json:"photo"`
}
