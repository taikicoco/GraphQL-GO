package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
)

type CountryRepository struct{}

func (ar *CountryRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Country, error) {
	countrys := []*model.Country{}
	err := db.SelectContext(ctx, &countrys, "SELECT country_id, name, img_url FROM countries")

	if err != nil {
		return nil, err
	}
	return countrys, nil
}

func (ar *CountryRepository) GetByID(ctx context.Context, db *sqlx.DB, countryID int) (*model.Country, error) {
	country := &model.Country{}
	err := db.Get(country, `select country_id, name, img_url from countries where country_id = ?`, countryID)
	if err != nil {
		return nil, err
	}
	return country, nil
}
