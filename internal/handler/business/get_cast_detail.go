package business

import (
	"context"
	"nitflex/internal/repository"
)

func (b *business) GetCastDetail(ctx context.Context, tmdbid int32) (*repository.CastInfo, error) {
	cast, err := b.repo.GetCastByID(ctx, tmdbid)
	if err != nil {
		return nil, err
	}

	cast.ProfilePath = "https://image.tmdb.org/t/p/w500" + cast.ProfilePath

	for i := range cast.MovieCredit.CastMovie {
		cast.MovieCredit.CastMovie[i].PosterPath = "https://image.tmdb.org/t/p/w500" + cast.MovieCredit.CastMovie[i].PosterPath
	}

	return cast, nil
}
