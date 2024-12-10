package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nitflex/internal/handler/models"
	"os"
)

func (t *tmdbAdapter) SearchMovies(ctx context.Context, request *SearchMoviesRequest) (*models.GetMoviesResponse, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%s&include_adult=false&language=%s&page=%d", request.Query, request.Language, request.Page)

	req, _ := http.NewRequest("GET", url, nil)

	access_token := os.Getenv("TMDB_ACCESS_TOKEN")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+access_token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response *models.GetMoviesResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
