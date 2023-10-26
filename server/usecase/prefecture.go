package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Prefecture struct {
	db        *sqlx.DB
	prefectureRepo *repository.PrefectureRepository
}

func NewPrefecture(db *sqlx.DB) *Prefecture {
	return &Prefecture{
		db: db,
	}
}

func (a *Prefecture) PrefectureList(ctx context.Context) ([]*model.Prefecture, error) {
	prefectures, err := a.prefectureRepo.GetAll(ctx, a.db)
	if err != nil {
		return nil, err
	}
	return prefectures, nil
}

func (a *Prefecture) GetPrefectureByID(ctx context.Context, prefectureID int) (*model.Prefecture, error) {
	prefecture, err := a.prefectureRepo.GetByID(ctx, a.db, prefectureID)
	if err != nil {
		return nil, err
	}
	return prefecture, nil
}
