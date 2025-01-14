package business

import (
	"context"

	"nitflex/internal/repository"
)

// Change return type to []*repository.Movie
func (b *business) SearchMoviesLLM(ctx context.Context, description string) ([]*repository.Movie, error) {
	// Call LLM adapter to get movie suggestions
	// llmResults, err := b.llmAdapter.SearchMoviesLLM(ctx, description)
	// if err != nil {
	// 	return nil, err
	// }

	// if len(llmResults) == 0 {
	// 	return []*repository.Movie{}, nil
	// }

	// // Parse JSON string into string array
	// var movieTitles []string
	// if err := json.Unmarshal([]byte(llmResults[0]), &movieTitles); err != nil {
	// 	return nil, err
	// }

	// // Search for movies in MongoDB using the titles
	// var allMovies []*repository.Movie
	// for _, title := range movieTitles {
	// 	movies, err := b.repo.SearchMoviesByQuery(ctx, title)
	// 	if err != nil {
	// 		continue // Skip if error occurs for one title
	// 	}
	// 	allMovies = append(allMovies, movies...)
	// }

	// return allMovies, nil
	return nil, nil
}
