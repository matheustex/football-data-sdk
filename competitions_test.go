package football

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompetitionService_Find(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/competitions/1", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"id": 2003,
			"area": {
				"id": 2163,
				"name": "Netherlands"
			},
			"name": "Eredivisie",			
			"plan": "TIER_ONE",
			"currentSeason": {
				"id": 4,
				"startDate": "2017-08-11T19:00:00Z",
				"endDate": "2018-05-20T16:45:00Z",
				"currentMatchday": 34
			},
			"seasons": [
				{
					"id": 4,
					"startDate": "2017-08-11T19:00:00Z",
					"endDate": "2018-05-20T16:45:00Z",
					"currentMatchday": 34
				}
			],
			"lastUpdated": "2018-06-05T00:17:50Z"
		}`)
	})

	season := Season{
		ID:              4,
		StartDate:       "2017-08-11T19:00:00Z",
		EndDate:         "2018-05-20T16:45:00Z",
		CurrentMatchday: 34,
	}

	expected := &Competition{
		ID: 2003,
		Area: Area{
			ID:   2163,
			Name: "Netherlands",
		},
		Name:          "Eredivisie",
		Code:          0,
		Plan:          "TIER_ONE",
		CurrentSeason: season,
		Seasons:       []Season{season},
		LastUpdated:   "2018-06-05T00:17:50Z",
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	competition, err := client.Competitions.Find(ctx, "1")

	assert.Nil(t, err)
	assert.Equal(t, expected, competition)
}
