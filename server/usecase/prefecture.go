package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Prefecture struct {
	db             *sqlx.DB
	prefectureRepo *repository.PrefectureRepository
}

func NewPrefecture(db *sqlx.DB) *Prefecture {
	return &Prefecture{
		db: db,
	}
}

func (p *Prefecture) GetPrefectures(ctx context.Context) ([]*model.Prefecture, error) {
	prefectures, err := p.prefectureRepo.GetAll(ctx, p.db)
	if err != nil {
		return nil, err
	}
	return prefectures, nil
}

func (p *Prefecture) GetPrefectureByID(ctx context.Context, prefectureID int) (*model.Prefecture, error) {
	prefecture, err := p.prefectureRepo.GetByID(ctx, p.db, prefectureID)
	if err != nil {
		return nil, err
	}
	return prefecture, nil
}

func (p *Prefecture) GetPrefectureByAnimeID(ctx context.Context, animeId int) ([]*model.Prefecture, error) {
	prefecture, err := p.prefectureRepo.GetByAnimeID(ctx, p.db, animeId)
	if err != nil {
		return nil, err
	}
	return prefecture, nil
}
