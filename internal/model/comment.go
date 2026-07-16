package model

import "time"

type Comment struct {
	ID        int64     `json:"id"`
	AnimeID   int64     `json:"animeId"`
	UserID    int64     `json:"userId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
