package models

type Movie struct {
	BackdropPath     string  `json:"backdrop_path"`
	Id               int     `json:"id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	MediaType        string  `json:"media_type"`
	Adult            bool    `json:"adult"`
	OriginalLanguage string  `json:"original_language"`
	GenreIds         []int   `json:"genre_ids"`
	Popularity       float64 `json:"popularity"`
	ReleaseDate      string  `json:"release_date"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type GetMoviesResponse struct {
	Page         int      `json:"page"`
	Results      []*Movie `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}
