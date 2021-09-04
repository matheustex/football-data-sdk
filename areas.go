package football

import (
	"context"
	"errors"
	"fmt"
)

type AreaService service

type Area struct {
	ID           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	CountryCode  string  `json:"countryCode,omitempty"`
	EnsignUrl    string  `json:"ensignUrl,omitempty"`
	ParentAreaID int     `json:"parentAreaId,omitempty"`
	ParentArea   string  `json:"parentArea,omitempty"`
	ChildAreas   *[]Area `json:"childAreas,omitempty"`
}

type AreaList struct {
	Count   int                    `json:"count,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Areas   []Area                 `json:"areas,omitempty"`
}

func (s *AreaService) Find(ctx context.Context, id string) (*Area, error) {
	if len(id) == 0 {
		return nil, errors.New("playerId is required")
	}

	area := &Area{}

	_, err := s.client.Get(fmt.Sprintf("areas/%s", id), nil, &area)
	if err != nil {
		return nil, err
	}

	return area, nil
}

func (s *AreaService) List(ctx context.Context) (*AreaList, error) {
	areas := &AreaList{}

	_, err := s.client.Get("areas", nil, &areas)
	if err != nil {
		return nil, err
	}

	return areas, nil
}
