package model

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type AnimeCategory struct {
	AnimeID    int64 `json:"animeId"`
	CategoryID int64 `json:"categoryId"`
}
