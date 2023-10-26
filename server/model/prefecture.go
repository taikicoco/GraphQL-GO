package model

type Prefecture struct {
	PrefectureID int    `db:"prefecture_id"`
	Name         string `db:"name"`
}
