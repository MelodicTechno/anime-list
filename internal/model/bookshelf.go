package model

import "time"

type Bookshelf struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type BookshelfItem struct {
	ID          int64 `json:"id"`
	BookshelfID int64 `json:"bookshelfId"`
	AnimeID     int64 `json:"animeId"`
}
