package repository

import (
	"time"
)

// /////// USERS /////////
type User struct {
	Id                    string     `bson:"_id,omitempty"`
	Username              string     `bson:"username"`
	Email                 string     `bson:"email"`
	Password              string     `bson:"password"`
	RefreshToken          string     `bson:"refresh_token,omitempty"`
	RefreshTokenExpiresAt *time.Time `bson:"refresh_token_expires_at,omitempty"`
	CreatedAt             time.Time  `bson:"created_at"`
	UpdatedAt             time.Time  `bson:"updated_at"`
	Watchlist             []int      `bson:"watchlist,omitempty"`
	FavoriteList          []int      `bson:"favorite_list,omitempty"`
}

func (User) TableName() string { return "users" }

// /////// MOVIES /////////
type Movie struct {
	Id                  string     `bson:"_id,omitempty"`
	TmdbId              int        `bson:"tmdb_id"`
	Adult               bool       `bson:"adult"`
	BackdropPath        string     `bson:"backdrop_path"`
	BelongsToCollection any        `bson:"belongs_to_collection"`
	Budget              int        `bson:"budget"`
	Categories          []string   `bson:"categories"`
	Genres              []Genre    `bson:"genres"`
	Homepage            string     `bson:"homepage"`
	ImdbId              string     `bson:"imdb_id"`
	OriginCountry       []string   `bson:"origin_country"`
	OriginalLanguage    string     `bson:"original_language"`
	OriginalTitle       string     `bson:"original_title"`
	Overview            string     `bson:"overview"`
	Popularity          float64    `bson:"popularity"`
	PosterPath          string     `bson:"poster_path"`
	ProductionCompanies []Company  `bson:"production_companies"`
	ProductionCountries []Country  `bson:"production_countries"`
	ReleaseDate         string     `bson:"release_date"`
	Revenue             int        `bson:"revenue"`
	Runtime             int        `bson:"runtime"`
	SpokenLanguages     []Language `bson:"spoken_languages"`
	Status              string     `bson:"status"`
	Tagline             string     `bson:"tagline"`
	Title               string     `bson:"title"`
	Video               bool       `bson:"video"`
	VoteAverage         float64    `bson:"vote_average"`
	VoteCount           int        `bson:"vote_count"`
	Credits             Credits    `bson:"credits"`
}

func (Movie) TableName() string { return "movies" }

type Genre struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

type Company struct {
	Id            int    `bson:"id"`
	LogoPath      string `bson:"logo_path"`
	Name          string `bson:"name"`
	OriginCountry string `bson:"origin_country"`
}

type Country struct {
	Iso3166_1 string `bson:"iso_3166_1"`
	Name      string `bson:"name"`
}

type Language struct {
	EnglishName string `bson:"english_name"`
	Iso639_1    string `bson:"iso_639_1"`
	Name        string `bson:"name"`
}

type Credits struct {
	Id   int    `bson:"id"`
	Cast []Cast `bson:"cast"`
}

type Cast struct {
	Adult              bool    `bson:"adult"`
	Gender             int     `bson:"gender"`
	Id                 int     `bson:"id"`
	KnownForDepartment string  `bson:"known_for_department"`
	Name               string  `bson:"name"`
	OriginalName       string  `bson:"original_name"`
	Popularity         float64 `bson:"popularity"`
	ProfilePath        string  `bson:"profile_path"`
	CastId             int     `bson:"cast_id"`
	Character          string  `bson:"character"`
	CreditId           string  `bson:"credit_id"`
	Order              int     `bson:"order"`
}

// /////// RATINGS /////////
type Rating struct {
	UserId  string `bson:"user_id"`
	MovieId string `bson:"movie_id"`
	Rating  int    `bson:"rating"`
}

// /////// CAST /////////
type CastInfo struct {
	Id           string      `bson:"_id"`
	TmdbId       int32       `bson:"tmdb_id"`
	Birthday     string      `bson:"birthday"`
	DeathDay     string      `bson:"deathday"`
	Name         string      `bson:"name"`
	PlaceOfBirth string      `bson:"place_of_birth"`
	Popularity   float64     `bson:"popularity"`
	Gender       int         `bson:"gender"`
	ProfilePath  string      `bson:"profile_path"`
	MovieCredit  MovieCredit `bson:"movie_credits"`
}

type MovieCredit struct {
	CastMovie []CastMovie `bson:"cast"`
}

type CastMovie struct {
	Id            int32  `bson:"id"`
	OriginalTitle string `bson:"original_title"`
	Title         string `bson:"title"`
	PosterPath    string `bson:"poster_path"`
	ReleaseDate   string `bson:"release_date"`
}

// /////// RECOMMENDATION /////////
type SimilarMovie struct {
	Id           string             `bson:"_id"`
	TmdbId       int32              `bson:"tmdb_id"`
	SimilarMovie []*SimilarMovieObj `bson:"similar_movies"`
}

type SimilarMovieObj struct {
	OriginalTitle string `bson:"original_title"`
	PosterPath    string `bson:"poster_path"`
	BackdropPath  string `bson:"backdrop_path"`
	ReleaseDate   string `bson:"release_date"`
	Id            int32  `bson:"id"`
}
