package football

import (
	"context"
	"errors"
	"fmt"
)

// AreaService provides methods for accessing information
// about matches.
type MatchService service

// Match represents data about a Match
type Match struct {
	ID            int             `json:"id,omitempty"`
	Competition   *Competition    `json:"competition,omitempty"`
	Season        *Season         `json:"season,omitempty"`
	UtcDate       string          `json:"utcDate,omitempty"`
	Status        string          `json:"status,omitempty"`
	Minute        interface{}     `json:"minute,omitempty"`
	Attendance    int             `json:"attendance,omitempty"`
	Venue         string          `json:"venue,omitempty"`
	Matchday      int             `json:"matchday,omitempty"`
	Stage         string          `json:"stage,omitempty"`
	Group         string          `json:"group,omitempty"`
	LastUpdated   string          `json:"lastUpdated,omitempty"`
	HomeTeam      *Team           `json:"homeTeam,omitempty"`
	AwayTeam      *Team           `json:"awayTeam,omitempty"`
	Score         *Score          `json:"score,omitempty"`
	Goals         []Goals         `json:"goals,omitempty"`
	Bookings      []Bookings      `json:"bookings,omitempty"`
	Substitutions []Substitutions `json:"substitutions,omitempty"`
	Referees      []Referees      `json:"referees,omitempty"`
}

type MatchesFiltersOptions struct {
	DateFrom     string `json:"dateFrom,omitempty"`
	DateTo       string `json:"dateTo,omitempty"`
	Status       string `json:"status,omitempty"`
	Competitions string `json:"competitions,omitempty"`
}

type MatchResponse struct {
	Head2Head Head2Head `json:"head2head,omitempty"`
	Match     Match     `json:"match,omitempty"`
}

// MatchesCompetition represents a collection of Matches for
// a competition
type MatchesCompetition struct {
	Count   int                    `json:"count,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Matches []Match                `json:"matches,omitempty"`
}

// Find takes a Match ID and returns the corresponding Match
// for that ID.
// https://www.football-data.org/documentation/api
func (s *MatchService) Find(ctx context.Context, id string) (*MatchResponse, error) {
	if len(id) == 0 {
		return nil, errors.New("Match ID is required")
	}

	match := &MatchResponse{}

	_, err := s.client.Get(fmt.Sprintf("matches/%s", id), nil, &match)
	if err != nil {
		return nil, err
	}

	return match, nil
}

// List returns a collection of all Matches.
// https://www.football-data.org/documentation/api
func (s *MatchService) List(ctx context.Context, filters *MatchesFiltersOptions) (*MatchesCompetition, error) {
	matchesCompetition := &MatchesCompetition{}

	_, err := s.client.Get("/matches", filters, &matchesCompetition)
	if err != nil {
		return nil, err
	}

	return matchesCompetition, nil
}
