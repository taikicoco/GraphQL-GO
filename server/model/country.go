package model

type Country struct {
	CountryID int    `db:"country_id"`
	Name      string `db:"name"`
	ImgURL    string `db:"img_url"`
}
