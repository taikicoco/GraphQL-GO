package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Gender struct {
	db         *sqlx.DB
	genderRepo *repository.GenderRepository
}

func NewGender(db *sqlx.DB) *Gender {
	return &Gender{
		db: db,
	}
}

func (g *Gender) GenderList(ctx context.Context) ([]*model.Gender, error) {
	genders, err := g.genderRepo.GetAll(ctx, g.db)
	if err != nil {
		return nil, err
	}
	return genders, nil
}

func (g *Gender) GetGenderByID(ctx context.Context, genderID int) (*model.Gender, error) {
	gender, err := g.genderRepo.GetByID(ctx, g.db, genderID)
	if err != nil {
		return nil, err
	}
	return gender, nil
}

func (g *Gender) GetGenderByGuideID(ctx context.Context, genderID int) (*model.Gender, error) {
	gender, err := g.genderRepo.GetByGuideID(ctx, g.db, genderID)
	if err != nil {
		return nil, err
	}
	return gender, nil
}
