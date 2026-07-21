package model

import "time"

type Anime struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	ReleaseDate  time.Time `json:"releaseDate"`
	Score        float64   `json:"score"`
	AiringStatus string    `json:"airingStatus"`
}
