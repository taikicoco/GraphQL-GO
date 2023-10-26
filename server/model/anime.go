package model

type Anime struct {
	AnimeID int    `db:"anime_id"`
	Name    string `db:"name"`
}
