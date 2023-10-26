package repository

import (
	"context"
	"server/model"

	"github.com/jmoiron/sqlx"
)

type SpotRepository struct{}

func (sr *SpotRepository) GetAll(ctx context.Context, db *sqlx.DB) ([]*model.Spot, error) {
	spots := []*model.Spot{}
	err := db.SelectContext(ctx, &spots,
		`SELECT spot_id, name, anime_img_url, real_img_url, address
		FROM spots`)

	if err != nil {
		return nil, err
	}
	return spots, nil
}

func (sr *SpotRepository) GetByID(ctx context.Context, db *sqlx.DB, spotID int) (*model.Spot, error) {
	spot := &model.Spot{}
	err := db.Get(spot,
		`select spot_id, name, anime_img_url, real_img_url, address
		from spots
		where spot_id = ?`, spotID)
	if err != nil {
		return nil, err
	}
	return spot, nil
}
