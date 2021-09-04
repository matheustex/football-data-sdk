package football

import (
	"context"
	"errors"
	"fmt"
)

type CompetitionService service

type Competition struct {
	ID            int      `json:"id,omitempty"`
	Area          Area     `json:"area,omitempty"`
	Name          string   `json:"name,omitempty"`
	Code          int      `json:"code,omitempty"`
	Plan          string   `json:"plan,omitempty"`
	CurrentSeason Season   `json:"currentSeason,omitempty"`
	Seasons       []Season `json:"seasons,omitempty"`
	LastUpdated   string   `json:"lastUpdated,omitempty"`
}

type CompetitionList struct {
	Count        int           `json:"id,omitempty"`
	Filters      interface{}   `json:"filters,omitempty"`
	Competitions []Competition `json:"competitions,omitempty"`
}

type CompetitionFiltersOptions struct {
	Areas string `json:"areas,omitempty"`
	Plan  string `json:"plan,omitempty"`
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

func (s *CompetitionService) List(ctx context.Context, filters CompetitionFiltersOptions) (*CompetitionList, error) {
	competitions := &CompetitionList{}

	_, err := s.client.Get("competitions", filters, &competitions)
	if err != nil {
		return nil, err
	}

	return competitions, nil
}
