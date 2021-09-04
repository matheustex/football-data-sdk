package football

import (
	"context"
	"errors"
	"fmt"
)

type TeamService service

type Team struct {
	ID                 int            `json:"id,omitempty"`
	Area               *Area          `json:"area,omitempty"`
	ActiveCompetitions *[]Competition `json:"activeCompetitions,omitempty"`
	Name               string         `json:"name,omitempty"`
	ShortName          string         `json:"shortName,omitempty"`
	Tla                string         `json:"tla,omitempty"`
	CrestURL           string         `json:"crestUrl,omitempty"`
	Address            string         `json:"address,omitempty"`
	Phone              string         `json:"phone,omitempty"`
	Website            string         `json:"website,omitempty"`
	Email              string         `json:"email,omitempty"`
	Founded            int            `json:"founded,omitempty"`
	ClubColors         string         `json:"clubColors,omitempty"`
	Venue              string         `json:"venue,omitempty"`
	Coach              *Coach         `json:"coach,omitempty"`
	Captain            *Player        `json:"captain,omitempty"`
	Squad              *[]Player      `json:"squad,omitempty"`
	Lineup             *[]Player      `json:"lineup,omitempty"`
	Bench              *[]Player      `json:"bench,omitempty"`
	LastUpdated        string         `json:"lastUpdated,omitempty"`
}

type TeamMatches struct {
	Count   int                    `json:"count,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Matches []Match                `json:"matches,omitempty"`
}

type TeamMatchesFiltersOptions struct {
	DateFrom string `json:"dateFrom,omitempty"`
	DateTo   string `json:"dateTo,omitempty"`
	Status   string `json:"status,omitempty"`
	Venue    string `json:"venue,omitempty"`
	Limit    int64  `json:"limit,omitempty"`
}

func (s *TeamService) Find(ctx context.Context, id string) (*Team, error) {
	if len(id) == 0 {
		return nil, errors.New("Team ID is required")
	}

	team := &Team{}

	_, err := s.client.Get(fmt.Sprintf("teams/%s", id), nil, &team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (s *TeamService) Matches(ctx context.Context, id string, filters *TeamMatchesFiltersOptions) (*TeamMatches, error) {
	if len(id) == 0 {
		return nil, errors.New("Team ID is required")
	}

	teamMatches := &TeamMatches{}

	_, err := s.client.Get(fmt.Sprintf("teams/%s/matches", id), filters, &teamMatches)
	if err != nil {
		return nil, err
	}

	return teamMatches, nil
}
