package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type AnimeRepository struct{}

func (ar *AnimeRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Anime, error) {
	animes := []*model.Anime{}
	err := db.SelectContext(ctx, &animes, "SELECT anime_id, name FROM animes")

	if err != nil {
		return nil, err
	}
	return animes, nil
}

func (ar *AnimeRepository) GetByID(ctx context.Context, db *sqlx.DB, animeID int) (*model.Anime, error) {
	anime := &model.Anime{}
	err := db.Get(anime, `select anime_id, name from animes where anime_id = ?`, animeID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return anime, nil
}
