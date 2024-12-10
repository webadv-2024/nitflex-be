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

func (t *tmdbAdapter) GetMovieDetail(ctx context.Context, request *GetMovieDetailRequest) (*models.GetMovieDetailResponse, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d", request.Id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	access_token := os.Getenv("TMDB_ACCESS_TOKEN")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+access_token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response *models.GetMovieDetailResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
