package model

type Spot struct {
	SpotID      int     `db:"spot_id"`
	Name        string  `db:"name"`
	AnimeImgURL *string `db:"anime_img_url"`
	RealImgURL  *string `db:"real_img_url"`
	Address     *string `db:"address"`
}
