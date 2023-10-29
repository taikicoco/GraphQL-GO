package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"server/graphql/generated"
	"server/graphql/generated/model"
)

// Prefecture is the resolver for the prefecture field.
func (r *animeResolver) Prefecture(ctx context.Context, obj *model.Anime) ([]*model.Prefecture, error) {
	res, err := r.prefecture.GetPrefectureByAnimeID(ctx, obj.AnimeID)
	if err != nil {
		return nil, err
	}

	prefectures := make([]*model.Prefecture, len(res))
	for i, v := range res {
		prefectures[i] = &model.Prefecture{
			PrefectureID: v.PrefectureID,
			AnimeID:      &obj.AnimeID,
			Name:         v.Name,
		}
	}
	return prefectures, nil
}

// Animes is the resolver for the animes field.
func (r *queryResolver) Animes(ctx context.Context) ([]*model.Anime, error) {
	res, err := r.anime.GetAnimes(ctx)
	if err != nil {
		return nil, err
	}

	animes := make([]*model.Anime, len(res))
	for i, v := range res {
		animes[i] = &model.Anime{
			AnimeID: v.AnimeID,
			Name:    v.Name,
			ImgURL:  &v.ImgURL,
		}
	}
	return animes, nil
}

// AnimesByAnimeIds is the resolver for the animesByAnimeIds field.
func (r *queryResolver) AnimesByAnimeIds(ctx context.Context, animeIds []int) ([]*model.Anime, error) {
	res, err := r.anime.GetAnimesByIDs(ctx, animeIds)
	if err != nil {
		return nil, err
	}

	animes := make([]*model.Anime, len(res))
	for i, v := range res {
		animes[i] = &model.Anime{
			AnimeID: v.AnimeID,
			Name:    v.Name,
			ImgURL:  &v.ImgURL,
		}
	}

	return animes, nil
}

// Anime is the resolver for the anime field.
func (r *queryResolver) Anime(ctx context.Context, animeID int) (*model.Anime, error) {
	res, err := r.anime.GetAnimeByID(ctx, animeID)
	if err != nil {
		return nil, err
	}

	anime := &model.Anime{
		AnimeID: res.AnimeID,
		Name:    res.Name,
		ImgURL:  &res.ImgURL,
	}

	return anime, nil
}

// Anime returns generated.AnimeResolver implementation.
func (r *Resolver) Anime() generated.AnimeResolver { return &animeResolver{r} }

type animeResolver struct{ *Resolver }
