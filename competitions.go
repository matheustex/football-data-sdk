package football

import (
	"context"
	"errors"
	"fmt"
)

// CompetitionService provides methods for accessing information
// about Competition.
type CompetitionService service

// Competition represents data about a Competition
type Competition struct {
	ID                       int      `json:"id,omitempty"`
	Area                     Area     `json:"area,omitempty"`
	Name                     string   `json:"name,omitempty"`
	Code                     string   `json:"code,omitempty"`
	NumberOfAvailableSeasons int      `json:"numberOfAvailableSeasons,omitempty"`
	Plan                     string   `json:"plan,omitempty"`
	CurrentSeason            Season   `json:"currentSeason,omitempty"`
	Seasons                  []Season `json:"seasons,omitempty"`
	LastUpdated              string   `json:"lastUpdated,omitempty"`
}

// CompetitionList represents a collection of Competitions
type CompetitionList struct {
	Count        int                    `json:"count,omitempty"`
	Filters      map[string]interface{} `json:"filters,omitempty"`
	Competitions []Competition          `json:"competitions,omitempty"`
}


type CompetitionFiltersOptions struct {
	Areas string `url:"areas,omitempty"`
	Plan  string `url:"plan,omitempty"`
}

// CompetitionTeams represents a collection of Teams for
// a competition
type CompetitionTeams struct {
	Count       int                    `json:"count,omitempty"`
	Filters     map[string]interface{} `json:"filters,omitempty"`
	Competition Competition            `json:"competition,omitempty"`
	Season      Season                 `json:"season,omitempty"`
	Teams       []Team                 `json:"teams,omitempty"`
}

type CompetitionTeamsFiltersOptions struct {
	Season string `url:"season,omitempty"`
	Stage  string `url:"stage,omitempty"`
}

// CompetitionTeams represents a collection of Standings for
// a competition
type CompetitionStandings struct {
	Filters     map[string]interface{} `json:"filters,omitempty"`
	Competition Competition            `json:"competition,omitempty"`
	Season      Season                 `json:"season,omitempty"`
	Standings   []Standing             `json:"standings,omitempty"`
}

type CompetitionStandingsFiltersOptions struct {
	StandingType StandingType `url:"standingType,omitempty"`
}

// CompetitionTeams represents a collection of Matches for
// a competition
type CompetitionMatches struct {
	Count       int                    `json:"count,omitempty"`
	Filters     map[string]interface{} `json:"filters,omitempty"`
	Competition Competition            `json:"competition,omitempty"`
	Matches     []Match                `json:"matches,omitempty"`
}

type CompetitionMatchesFiltersOptions struct {
	DateFrom string `json:"dateFrom,omitempty"`
	DateTo   string `json:"dateTo,omitempty"`
	Stage    string `json:"stage,omitempty"`
	Status   Status `json:"status,omitempty"`
	MatchDay string `json:"matchday,omitempty"`
	Group    string `json:"group,omitempty"`
	Season   string `json:"season,omitempty"`
}

// CompetitionTeams represents a collection of Scorers for
// a competition
type CompetitionScorers struct {
	Count       int                    `json:"count,omitempty"`
	Filters     map[string]interface{} `json:"filters,omitempty"`
	Competition Competition            `json:"competition,omitempty"`
	Season      Season                 `json:"season,omitempty"`
	Scorers     []Scorer               `json:"scorers,omitempty"`
}

type CompetitionScorersFiltersOptions struct {
	Limit string `json:"limit,omitempty"`
}

// Find takes a Competition ID and returns the corresponding Competition
// for that ID.
// https://www.football-data.org/documentation/api
func (s *CompetitionService) Find(ctx context.Context, id string) (*Competition, error) {
	if len(id) == 0 {
		return nil, errors.New("Competition ID is required")
	}

	competition := &Competition{}

	_, err := s.client.Get(fmt.Sprintf("competitions/%s", id), nil, &competition)
	if err != nil {
		return nil, err
	}

	return competition, nil
}

// List returns a collection of all competitions.
// https://www.football-data.org/documentation/api
func (s *CompetitionService) List(ctx context.Context, filters *CompetitionFiltersOptions) (*CompetitionList, error) {
	competitions := &CompetitionList{}

	_, err := s.client.Get("competitions", filters, &competitions)
	if err != nil {
		return nil, err
	}

	return competitions, nil
}

// Teams takes a Competition ID and returns a 
// collection of all teams for that competition.
// https://www.football-data.org/documentation/api
func (s *CompetitionService) Teams(ctx context.Context, id string, filters *CompetitionTeamsFiltersOptions) (*CompetitionTeams, error) {
	competitionTeams := &CompetitionTeams{}

	if len(id) == 0 {
		return nil, errors.New("Competition ID is required")
	}
	_, err := s.client.Get(fmt.Sprintf("competitions/%s/teams", id), filters, &competitionTeams)
	if err != nil {
		return nil, err
	}

	return competitionTeams, nil
}

// Teams takes a Competition ID and returns a 
// collection of all standings for that competition.
// https://www.football-data.org/documentation/api
func (s *CompetitionService) Standings(ctx context.Context, id string, filters *CompetitionStandingsFiltersOptions) (*CompetitionStandings, error) {
	competitionStandings := &CompetitionStandings{}

	if len(id) == 0 {
		return nil, errors.New("Competition ID is required")
	}
	_, err := s.client.Get(fmt.Sprintf("competitions/%s/standings", id), filters, &competitionStandings)
	if err != nil {
		return nil, err
	}

	return competitionStandings, nil
}

// Teams takes a Competition ID and returns a 
// collection of all matches for that competition.
// https://www.football-data.org/documentation/api
func (s *CompetitionService) Matches(ctx context.Context, id string, filters *CompetitionMatchesFiltersOptions) (*CompetitionMatches, error) {
	if len(id) == 0 {
		return nil, errors.New("Competition ID is required")
	}

	competitionMatches := &CompetitionMatches{}

	_, err := s.client.Get(fmt.Sprintf("competitions/%s/matches", id), filters, &competitionMatches)
	if err != nil {
		return nil, err
	}

	return competitionMatches, nil
}

// Teams takes a Competition ID and returns a 
// collection of all scorers for that competition.
// https://www.football-data.org/documentation/api
func (s *CompetitionService) Scorers(ctx context.Context, id string, filters *CompetitionScorersFiltersOptions) (*CompetitionScorers, error) {
	if len(id) == 0 {
		return nil, errors.New("Competition ID is required")
	}

	competitionScorers := &CompetitionScorers{}

	_, err := s.client.Get(fmt.Sprintf("competitions/%s/scorers", id), filters, &competitionScorers)
	if err != nil {
		return nil, err
	}

	return competitionScorers, nil
}
