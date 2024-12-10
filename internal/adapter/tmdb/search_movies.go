package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nitflex/internal/handler/models"
)

func (t *tmdbAdapter) SearchMovies(ctx context.Context, request *SearchMoviesRequest) (*models.GetMoviesResponse, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%s&include_adult=false&language=%s&page=%d", request.Query, request.Language, request.Page)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI1NjVmNzdjOTgwNmVjOWZhYjI4ZDhkZTJhNTI1NzcyOCIsIm5iZiI6MTczMzc1NTU4MS4wNjksInN1YiI6IjY3NTcwMmJkNmYzN2Y2ZTg5YWRjODBkYSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.JEpxWGOdU2Y3LehfnmzQQlrXPAZfYTrnIoRc2cKZCT0")

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
