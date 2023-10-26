package model

type Spot struct {
	SpotID      int     `json:"spotId" db:"spot_id"`
	Name        string  `json:"name" db:"name"`
	AnimeImgURL *string `json:"animeImgUrl" db:"anime_img_url"`
	RealImgURL  *string `json:"realImgUrl" db:"real_img_url"`
	Address     *string `json:"address" db:"address"`
}
