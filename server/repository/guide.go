package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
)

type GuideRepository struct{}

func (gr *GuideRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Guide, error) {
	guides := []*model.Guide{}
	err := db.SelectContext(ctx, &guides,
		`select guide_id, name, age, comment, stance, favorite_character, gender_id, country_id
		from guides`)

	if err != nil {
		return nil, err
	}
	return guides, nil
}

func (gr *GuideRepository) GetByID(ctx context.Context, db *sqlx.DB, guideID int) (*model.Guide, error) {
	guide := &model.Guide{}
	err := db.Get(guide,
		`select guide_id, name, age, comment, stance, favorite_character, gender_id, country_id
		from guides
		where guide_id = ?`, guideID)
	if err != nil {
		return nil, err
	}
	return guide, nil
}
