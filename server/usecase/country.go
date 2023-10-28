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

func (c *Country) CountryList(ctx context.Context) ([]*model.Country, error) {
	countrys, err := c.countryRepo.GetAll(ctx, c.db)
	if err != nil {
		return nil, err
	}
	return countrys, nil
}

func (c *Country) GetCountryByID(ctx context.Context, countryID int) (*model.Country, error) {
	country, err := c.countryRepo.GetByID(ctx, c.db, countryID)
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (c *Country) GetCountryByGuideID(ctx context.Context, guideID int) (*model.Country, error) {
	country, err := c.countryRepo.GetByGuideID(ctx, c.db, guideID)
	if err != nil {
		return nil, err
	}
	return country, nil
}
