package football

import (
	"context"
	"errors"
	"fmt"
)

// AreaService provides methods for accessing information
// about areas.
type AreaService service

// Area represents data about a Area.
type Area struct {
	ID           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	CountryCode  string  `json:"countryCode,omitempty"`
	EnsignUrl    string  `json:"ensignUrl,omitempty"`
	ParentAreaID int     `json:"parentAreaId,omitempty"`
	ParentArea   string  `json:"parentArea,omitempty"`
	ChildAreas   *[]Area `json:"childAreas,omitempty"`
}

// AreaList represents a collection of Areas
type AreaList struct {
	Count   int                    `json:"count,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Areas   []Area                 `json:"areas,omitempty"`
}

// Find takes a Area ID and returns the corresponding Area
// for that ID.
// https://www.football-data.org/documentation/api
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

// List returns a collection of all areas.
// https://www.football-data.org/documentation/api
func (s *AreaService) List(ctx context.Context) (*AreaList, error) {
	areas := &AreaList{}

	_, err := s.client.Get("areas", nil, &areas)
	if err != nil {
		return nil, err
	}

	return areas, nil
}
