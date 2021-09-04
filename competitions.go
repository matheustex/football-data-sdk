package football

import (
	"context"
	"errors"
	"fmt"
)

type CompetitionService service

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

type CompetitionList struct {
	Count        int                    `json:"count,omitempty"`
	Filters      map[string]interface{} `json:"filters,omitempty"`
	Competitions []Competition          `json:"competitions,omitempty"`
}

type CompetitionFiltersOptions struct {
	Areas string `url:"areas,omitempty"`
	Plan  string `url:"plan,omitempty"`
}

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

type CompetitionStandings struct {
	Filters     map[string]interface{} `json:"filters,omitempty"`
	Competition Competition            `json:"competition,omitempty"`
	Season      Season                 `json:"season,omitempty"`
	Standings   []Standing             `json:"standings,omitempty"`
}

type CompetitionStandingsFiltersOptions struct {
	StandingType StandingType `url:"standingType,omitempty"`
}

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

func (s *CompetitionService) List(ctx context.Context, filters *CompetitionFiltersOptions) (*CompetitionList, error) {
	competitions := &CompetitionList{}

	_, err := s.client.Get("competitions", filters, &competitions)
	if err != nil {
		return nil, err
	}

	return competitions, nil
}

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
