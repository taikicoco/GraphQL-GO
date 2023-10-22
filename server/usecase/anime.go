package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Anime struct {
	db        *sqlx.DB
	animeRepo *repository.AnimeRepository
}

func NewAnime(db *sqlx.DB) *Anime {
	return &Anime{
		db: db,
	}
}
func (a *Anime) AnimeList(ctx context.Context) ([]*model.Anime, error) {
	animes, err := a.animeRepo.GetAll(ctx, a.db)
	if err != nil {
		return nil, err
	}
	return animes, nil
}
