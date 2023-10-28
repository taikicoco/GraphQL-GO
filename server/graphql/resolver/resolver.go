package resolver

import (
	"server/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	anime      *usecase.Anime
	spot       *usecase.Spot
	prefecture *usecase.Prefecture
	gender     *usecase.Gender
	country    *usecase.Country
	guide      *usecase.Guide
}

func NewResolver(anime *usecase.Anime, spot *usecase.Spot, prefecture *usecase.Prefecture,
	gender *usecase.Gender, country *usecase.Country, guide *usecase.Guide) *Resolver {
	return &Resolver{
		anime:      anime,
		spot:       spot,
		prefecture: prefecture,
		gender:     gender,
		country:    country,
		guide:      guide,
	}
}
