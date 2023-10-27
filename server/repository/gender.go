package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
)

type GenderRepository struct{}

func (ar *GenderRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Gender, error) {
	genders := []*model.Gender{}
	err := db.SelectContext(ctx, &genders, "SELECT gender_id, gender FROM genders")

	if err != nil {
		return nil, err
	}
	return genders, nil
}

func (ar *GenderRepository) GetByID(ctx context.Context, db *sqlx.DB, genderID int) (*model.Gender, error) {
	gender := &model.Gender{}
	err := db.Get(gender, `select gender_id, gender from genders where gender_id = ?`, genderID)
	if err != nil {
		return nil, err
	}
	return gender, nil

}
