package usecase

import (
	"context"
	"server/model"
	"server/repository"

	"github.com/jmoiron/sqlx"
)

type Country struct {
	db          *sqlx.DB
	countryRepo *repository.CountryRepository
}

func NewCountry(db *sqlx.DB) *Country {
	return &Country{
		db: db,
	}
}

func (a *Country) CountryList(ctx context.Context) ([]*model.Country, error) {
	countrys, err := a.countryRepo.GetAll(ctx, a.db)
	if err != nil {
		return nil, err
	}
	return countrys, nil
}

func (a *Country) GetCountryByID(ctx context.Context, countryID int) (*model.Country, error) {
	country, err := a.countryRepo.GetByID(ctx, a.db, countryID)
	if err != nil {
		return nil, err
	}
	return country, nil
}
