package model

type Anime struct {
	AnimeID int    `json:"animeID" db:"anime_id"`
	Name    string `json:"name" db:"name"`
}
