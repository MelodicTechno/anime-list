package model

import "time"

type WatchPlan struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	AnimeID   int64     `json:"animeId"`
	Status    string    `json:"status"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
