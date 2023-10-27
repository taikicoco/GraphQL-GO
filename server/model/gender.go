package model

type Gender struct {
	GenderID int    `db:"gender_id"`
	Gender   string `db:"gender"`
}
