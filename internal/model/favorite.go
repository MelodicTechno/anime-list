package model

import "time"

type Favorite struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type FavoriteItem struct {
	ID         int64 `json:"id"`
	FavoriteID int64 `json:"favoriteId"`
	AnimeID    int64 `json:"animeId"`
}
