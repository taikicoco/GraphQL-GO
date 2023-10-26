package model

type Prefecture struct {
	PrefectureID int    `json:"prefectureId" db:"prefecture_id"`
	Name         string `json:"name" db:"name"`
}
