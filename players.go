package football

import (
	"context"
	"errors"
	"fmt"
)

type PlayerService service

type Player struct {
	ID             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	FirstName      string `json:"firstName,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	DateOfBirth    string `json:"dateOfBirth,omitempty"`
	CountryOfBirth string `json:"countryOfBirth,omitempty"`
	Nationality    string `json:"nationality,omitempty"`
	Position       string `json:"position,omitempty"`
	ShirtNumber    int    `json:"shirtNumber,omitempty"`
	LastUpdated    string `json:"lastUpdated,omitempty"`
	Role           string `json:"role,omitempty"`
}

type PlayerFiltersOptions struct {
	DateFrom     string `json:"dateFrom,omitempty"`
	DateTo       string `json:"dateTo,omitempty"`
	Status       string `json:"status,omitempty"`
	Competitions string `json:"competitions,omitempty"`
	Limit        int64  `json:"limit,omitempty"`
}

type PlayerMatches struct {
	Count   int                    `json:"count,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Player  Player                 `json:"player,omitempty"`
	Matches []Match                `json:"matches,omitempty"`
}

func (s *PlayerService) Find(ctx context.Context, id string) (*Player, error) {
	if len(id) == 0 {
		return nil, errors.New("playerId is required")
	}

	player := &Player{}

	_, err := s.client.Get(fmt.Sprintf("players/%s", id), nil, &player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func (s *PlayerService) Matches(ctx context.Context, id string, filters *PlayerFiltersOptions) (*PlayerMatches, error) {
	if len(id) == 0 {
		return nil, errors.New("playerId is required")
	}

	playerMatches := &PlayerMatches{}

	_, err := s.client.Get(fmt.Sprintf("players/%s/matches", id), filters, &playerMatches)
	if err != nil {
		return nil, err
	}

	return playerMatches, nil
}
