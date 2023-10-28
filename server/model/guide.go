package model

type Guide struct {
	GuideID           int    `db:"guide_id"`
	Name              string `db:"name"`
	Age               int    `db:"age"`
	Comment           string `db:"comment"`
	Stance            string `db:"stance"`
	FavoriteCharacter string `db:"favorite_character"`
	GenderID          int    `db:"gender_id"`
	CountryID         int    `db:"country_id"`
}
