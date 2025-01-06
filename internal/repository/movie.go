package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repository) GetMovieByID(ctx context.Context, movieID string) (*Movie, error) {
	var movie Movie

	// Convert movieID string to int
	movieIDInt, err := strconv.Atoi(movieID)
	if err != nil {
		return nil, fmt.Errorf("invalid movie ID format: %v", err)
	}

	err = r.mongodb.Collection(Movie{}.TableName()).
		FindOne(ctx, bson.M{"tmdb_id": movieIDInt}).
		Decode(&movie)

	return &movie, err
}

func (r *repository) GetMoviesList(ctx context.Context, movieIDs []int) ([]*Movie, error) {
	var movies []*Movie

	collection := r.mongodb.Collection(Movie{}.TableName())

	// Create filter for multiple movie IDs
	filter := bson.M{"tmdb_id": bson.M{"$in": movieIDs}}

	// Find all matching movies
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding movies: %v", err)
	}
	defer cursor.Close(ctx)

	// Decode all movies into the slice
	if err := cursor.All(ctx, &movies); err != nil {
		return nil, fmt.Errorf("error decoding movies: %v", err)
	}

	return movies, nil
}

func (r *repository) GetTrendingMoviesInDay(ctx context.Context) ([]*Movie, error) {
	var movies []*Movie

	collection := r.mongodb.Collection("movies_trending_day")
	options := options.Find().SetLimit(10)

	cursor, err := collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, fmt.Errorf("error finding movies: %v", err)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &movies); err != nil {
		return nil, fmt.Errorf("error decoding movies: %v", err)
	}

	return movies, nil
}

func (r *repository) SearchMoviesByQuery(ctx context.Context, title string) ([]*Movie, error) {
	var movies []*Movie

	collection := r.mongodb.Collection("movies")

	// Create filter for title
	filter := bson.M{"title": bson.M{"$regex": title, "$options": "i"}} // Case-insensitive search

	// Find all matching movies
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding movies: %v", err)
	}
	defer cursor.Close(ctx)

	// Decode all movies into the slice
	if err := cursor.All(ctx, &movies); err != nil {
		return nil, fmt.Errorf("error decoding movies: %v", err)
	}

	return movies, nil
}

func (r *repository) FilterMovies(ctx context.Context, params *FilterMoviesParams) ([]*Movie, error) {
	var movies []*Movie
	collection := r.mongodb.Collection("movies")

	filter := bson.M{}

	// Genre filter
	if params.Genres != "" {
		genreIDs := strings.Split(params.Genres, ",")
		intGenreIDs := make([]int, len(genreIDs))
		for i, id := range genreIDs {
			intGenreIDs[i] = cast.ToInt(id)
		}
		filter["genres.id"] = bson.M{"$in": intGenreIDs}
	}

	// Rating filters
	if params.MinRating > 0 || params.MaxRating > 0 {
		ratingFilter := bson.M{}
		if params.MinRating > 0 {
			ratingFilter["$gte"] = params.MinRating
		}
		if params.MaxRating > 0 {
			ratingFilter["$lte"] = params.MaxRating
		}
		filter["vote_average"] = ratingFilter
	}

	// Release date filters
	if params.ReleaseDateGte != "" || params.ReleaseDateLte != "" {
		dateFilter := bson.M{}

		if params.ReleaseDateGte != "" {
			fromDate, err := time.Parse("2006-01-02", params.ReleaseDateGte)
			if err == nil {
				dateFilter["$gte"] = fromDate.Format("2006-01-02")
			} else {
				fmt.Printf("Error parsing from date: %v\n", err)
			}
		}

		if params.ReleaseDateLte != "" {
			toDate, err := time.Parse("2006-01-02", params.ReleaseDateLte)
			if err == nil {
				dateFilter["$lte"] = toDate.Format("2006-01-02")
			} else {
				fmt.Printf("Error parsing to date: %v\n", err)
			}
		}

		if len(dateFilter) > 0 {
			filter["release_date"] = dateFilter
		}
	}

	// Actor filter
	if params.Actors != "" {
		actorNames := strings.Split(params.Actors, ",")
		// Trim whitespace from actor names
		for i, name := range actorNames {
			actorNames[i] = strings.TrimSpace(name)
		}
		filter["credits.cast.name"] = bson.M{"$in": actorNames}
	}

	fmt.Printf("Final MongoDB filter: %+v\n", filter)

	// Find all matching movies
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding movies: %v", err)
	}
	defer cursor.Close(ctx)

	// Decode all movies into the slice
	if err := cursor.All(ctx, &movies); err != nil {
		return nil, fmt.Errorf("error decoding movies: %v", err)
	}

	// Check if any movies were found
	if len(movies) == 0 {
		fmt.Println("No documents found matching the filter")
		return movies, nil
	}

	fmt.Printf("Found %d movies matching the criteria\n", len(movies))
	return movies, nil
}
