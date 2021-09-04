package football

import (
	"context"
	"errors"
	"fmt"
)

type PlayersService service

type Player struct {
	ID             *int64  `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
	FirstName      *string `json:"firstName,omitempty"`
	LastName       *string `json:"lastName,omitempty"`
	DateOfBirth    *string `json:"dateOfBirth,omitempty"`
	CountryOfBirth *string `json:"countryOfBirth,omitempty"`
	Nationality    *string `json:"nationality,omitempty"`
	Position       *string `json:"position,omitempty"`
	LastUpdated    *string `json:"lastUpdated,omitempty"`
}

type PlayerFiltersOptions struct {
	DateFrom     *string `json:"dateFrom,omitempty"`
	DateTo       *string `json:"dateTo,omitempty"`
	Status       *string `json:"status,omitempty"`
	Competitions *string `json:"competitions,omitempty"`
	Limit        *int64  `json:"limit,omitempty"`
}

func (s *PlayersService) Find(ctx context.Context, id string) (*Player, error) {
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

func (s *PlayersService) FindMatches(ctx context.Context, id string) (*Player, error) {
	if len(id) == 0 {
		return nil, errors.New("playerId is required")
	}

	player := &Player{}

	_, err := s.client.Get(fmt.Sprintf("players/%s/matches", id), nil, &player)
	if err != nil {
		return nil, err
	}

	return player, nil
}
