package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Gender struct {
	db        *sqlx.DB
	genderRepo *repository.GenderRepository
}

func NewGender(db *sqlx.DB) *Gender {
	return &Gender{
		db: db,
	}
}

func (a *Gender) GenderList(ctx context.Context) ([]*model.Gender, error) {
	genders, err := a.genderRepo.GetAll(ctx, a.db)
	if err != nil {
		return nil, err
	}
	return genders, nil
}

func (a *Gender) GetGenderByID(ctx context.Context, genderID int) (*model.Gender, error) {
	gender, err := a.genderRepo.GetByID(ctx, a.db, genderID)
	if err != nil {
		return nil, err
	}
	return gender, nil
}
