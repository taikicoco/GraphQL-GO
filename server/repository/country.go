package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
)

type CountryRepository struct{}

func (cr *CountryRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Country, error) {
	countrys := []*model.Country{}
	err := db.SelectContext(ctx, &countrys, "SELECT country_id, name, img_url FROM countries")

	if err != nil {
		return nil, err
	}
	return countrys, nil
}

func (cr *CountryRepository) GetByID(ctx context.Context, db *sqlx.DB, countryID int) (*model.Country, error) {
	country := &model.Country{}
	err := db.Get(country, `select country_id, name, img_url from countries where country_id = ?`, countryID)
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (cr *CountryRepository) GetByGuideID(ctx context.Context, db *sqlx.DB, guideID int) (*model.Country, error) {
	country := &model.Country{}
	err := db.Get(country,
		`select c.country_id, c.name, c.img_url
		from countries c
		inner join guides g on g.country_id = c.country_id
		where g.guide_id = ?`, guideID)
	if err != nil {
		return nil, err
	}
	return country, nil
}
