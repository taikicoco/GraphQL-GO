package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
)

type AnimeRepository struct{}

func (ar *AnimeRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Anime, error) {
	animes := []*model.Anime{}
	err := db.SelectContext(ctx, &animes,
		"SELECT anime_id, name FROM animes")

	if err != nil {
		return nil, err
	}
	return animes, nil
}
