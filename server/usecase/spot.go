package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Spot struct {
	db       *sqlx.DB
	spotRepo *repository.SpotRepository
}

func NewSpot(db *sqlx.DB) *Spot {
	return &Spot{
		db: db,
	}
}

func (s *Spot) SpotList(ctx context.Context) ([]*model.Spot, error) {
	spots, err := s.spotRepo.GetAll(ctx, s.db)
	if err != nil {
		return nil, err
	}
	return spots, nil
}

func (s *Spot) GetSpotByID(ctx context.Context, spotID int) (*model.Spot, error) {
	spot, err := s.spotRepo.GetByID(ctx, s.db, spotID)
	if err != nil {
		return nil, err
	}
	return spot, nil
}

func (s *Spot) GetSpotByPrefectureID(ctx context.Context, animeId int, prefectureId int) ([]*model.Spot, error) {
	spots, err := s.spotRepo.GetByPrefectureID(ctx, s.db, animeId, prefectureId)
	if err != nil {
		return nil, err
	}
	return spots, nil
}
