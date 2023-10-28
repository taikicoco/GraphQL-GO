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

func (gr *GuideRepository) GetByAnimesIDsAndPrefectureIDs(ctx context.Context, db *sqlx.DB, animeIDs []int, prefectureIDs []int) ([]*model.Guide, error) {
	guides := []*model.Guide{}
	query := `
		select g.guide_id, name, age, comment, stance, favorite_character, gender_id, country_id
		from guides g
		inner join guide_animes ga on g.guide_id = ga.guide_id
		inner join guide_prefectures gp on g.guide_id = gp.guide_id
		where ga.anime_id in (?)
		and gp.prefecture_id in (?)
	`
	query, args, err := sqlx.In(query, animeIDs, prefectureIDs)
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)

	err = db.Select(&guides, query, args...)
	if err != nil {
		return nil, err
	}
	return guides, nil
}
