package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Spot struct {
	db        *sqlx.DB
	spotRepo *repository.SpotRepository
}

func NewSpot(db *sqlx.DB) *Spot {
	return &Spot{
		db: db,
	}
}

func (a *Spot) SpotList(ctx context.Context) ([]*model.Spot, error) {
	spots, err := a.spotRepo.GetAll(ctx, a.db)
	if err != nil {
		return nil, err
	}
	return spots, nil
}

func (a *Spot) GetSpotByID(ctx context.Context, spotID int) (*model.Spot, error) {
	spot, err := a.spotRepo.GetByID(ctx, a.db, spotID)
	if err != nil {
		return nil, err
	}
	return spot, nil
}
