package repository

import (
	"context"
	"server/model"
	// "github.com/jmoiron/sqlx"
)

type AnimeRepository struct{}

func (ar *AnimeRepository) GetAll(ctx context.Context) ([]*model.Anime, error) {
	animes := []*model.Anime{
		{
			AnimeID: 1,
			Name:    "anime1",
		},
		{
			AnimeID: 2,
			Name:    "anime2",
		},
	}
	return animes, nil

}
