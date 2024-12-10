package adapter

type GetTrendingMoviesRequest struct {
	TimeWindow string
	Language   string
}

type SearchMoviesRequest struct {
	Query    string
	Language string
	Page     int
}

type GetMovieDetailRequest struct {
	Id  int
}