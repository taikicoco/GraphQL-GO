package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
)

type PrefectureRepository struct{}

func (pr *PrefectureRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Prefecture, error) {
	prefectures := []*model.Prefecture{}
	err := db.SelectContext(ctx, &prefectures, "SELECT prefecture_id, name FROM prefectures")

	if err != nil {
		return nil, err
	}
	return prefectures, nil
}

func (pr *PrefectureRepository) GetByID(ctx context.Context, db *sqlx.DB, prefectureID int) (*model.Prefecture, error) {
	prefecture := &model.Prefecture{}
	err := db.Get(prefecture, `select prefecture_id, name from prefectures where prefecture_id = ?`, prefectureID)
	if err != nil {
		return nil, err
	}
	return prefecture, nil
}

func (pr *PrefectureRepository) GetByAnimeID(ctx context.Context, db *sqlx.DB, animeID int) ([]*model.Prefecture, error) {
	prefecture := []*model.Prefecture{}
	err := db.Select(&prefecture,
		`select distinct p.prefecture_id, p.name
		from prefectures p
		inner join spots s on p.prefecture_id = s.prefecture_id
		where s.anime_id = ?`, animeID)
	if err != nil {
		return nil, err
	}


	return prefecture, nil
}
