package models

type Movie struct {
	BackdropPath     string  `bson:"backdrop_path"`
	Id               int     `bson:"id"`
	Title            string  `bson:"title"`
	OriginalTitle    string  `bson:"original_title"`
	Overview         string  `bson:"overview"`
	PosterPath       string  `bson:"poster_path"`
	MediaType        string  `bson:"media_type"`
	Adult            bool    `bson:"adult"`
	OriginalLanguage string  `bson:"original_language"`
	GenreIds         []int   `bson:"genre_ids"`
	Popularity       float64 `bson:"popularity"`
	ReleaseDate      string  `bson:"release_date"`
	Video            bool    `bson:"video"`
	VoteAverage      float64 `bson:"vote_average"`
	VoteCount        int     `bson:"vote_count"`
}

type GetMoviesResponse struct {
	Page         int      `bson:"page"`
	Results      []*Movie `bson:"results"`
	TotalPages   int      `bson:"total_pages"`
	TotalResults int      `bson:"total_results"`
}

type TextResponse struct {
	Message string `bson:"message"`
}

type MovieListResponse struct {
	Results []*Movie `bson:"results"`
}
