package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Guide struct {
	db        *sqlx.DB
	guideRepo *repository.GuideRepository
}

func NewGuide(db *sqlx.DB) *Guide {
	return &Guide{
		db: db,
	}
}

func (g *Guide) GetGuides(ctx context.Context) ([]*model.Guide, error) {
	guides, err := g.guideRepo.GetAll(ctx, g.db)
	if err != nil {
		return nil, err
	}
	return guides, nil
}

func (g *Guide) GetGuideByID(ctx context.Context, guideID int) (*model.Guide, error) {
	guide, err := g.guideRepo.GetByID(ctx, g.db, guideID)
	if err != nil {
		return nil, err
	}
	return guide, nil
}

func (g *Guide) GetGuidesByAnimeIDsAndPrefectureIDs(ctx context.Context, animeIDs []int, prefectureIDs []int) ([]*model.Guide, error) {
	guides, err := g.guideRepo.GetByAnimesIDsAndPrefectureIDs(ctx, g.db, animeIDs, prefectureIDs)
	if err != nil {
		return nil, err
	}
	return guides, nil
}
