package models

import "nitflex/internal/repository"

type RatingsResponse struct {
	Results []*repository.Rating `json:"results"`
}
