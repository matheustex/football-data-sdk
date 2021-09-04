package football

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamService_Find(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/teams/1", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"id": 18,
			"area": {
				"id": 2088,
				"name": "Germany"
			},
			"activeCompetitions": [
				{
					"id": 2002,
					"area": {
						"id": 2088,
						"name": "Germany"
					},
					"name": "Bundesliga",
					"code": "BL1",
					"plan": "TIER_ONE",
					"lastUpdated": "2021-04-17T00:20:14Z"
				}				
			],
			"name": "Borussia Mönchengladbach",
			"shortName": "M'gladbach",
			"tla": "BMG",
			"crestUrl": "https://crests.football-data.org/18.svg",
			"clubColors": "Black / White / Green",
			"venue": "Stadion im Borussia-Park",
			"squad": [
				{
					"id": 3176,
					"name": "Matthias Ginter",
					"position": "Defender",
					"dateOfBirth": "1994-01-19T00:00:00Z",
					"countryOfBirth": "Germany",
					"nationality": "Germany",
					"shirtNumber": null,
					"role": "PLAYER"
				}
			],
			"lastUpdated": "2020-11-26T02:04:29Z"
		}`)
	})

	expected := Team{
		ID: 18,
		Area: &Area{
			ID:   2088,
			Name: "Germany",
		},
		ActiveCompetitions: &[]Competition{
			{
				ID: 2002,
				Area: Area{
					ID:   2088,
					Name: "Germany",
				},
				Name:        "Bundesliga",
				Code:        "BL1",
				Plan:        "TIER_ONE",
				LastUpdated: "2021-04-17T00:20:14Z",
			},
		},
		Name:       "Borussia Mönchengladbach",
		ShortName:  "M'gladbach",
		Tla:        "BMG",
		CrestURL:   "https://crests.football-data.org/18.svg",
		ClubColors: "Black / White / Green",
		Venue:      "Stadion im Borussia-Park",
		Squad: &[]Player{
			{
				ID:             3176,
				Name:           "Matthias Ginter",
				Position:       "Defender",
				DateOfBirth:    "1994-01-19T00:00:00Z",
				CountryOfBirth: "Germany",
				Nationality:    "Germany",
				Role:           "PLAYER",
			},
		},
		LastUpdated: "2020-11-26T02:04:29Z",
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	matchResponse, err := client.Teams.Find(ctx, "1")

	assert.Nil(t, err)
	assert.Equal(t, expected, *matchResponse)
}

func TestTeamService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/teams/18/matches", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"count": 1,
			"matches": [
				{
					"id": 328846,
					"season": {
						"id": 734,
						"startDate": "2021-06-26",
						"endDate": "2022-05-22",
						"currentMatchday": 1
					},
					"utcDate": "2021-06-22T18:00:00Z",
					"status": "FINISHED",
					"matchday": null,
					"stage": "PRELIMINARY_ROUND",
					"group": null,
					"lastUpdated": "2021-09-04T16:20:05Z",
					"score": {
						"winner": "AWAY_TEAM",
						"duration": "REGULAR",
						"fullTime": {
							"homeTeam": 0,
							"awayTeam": 1
						},
						"halfTime": {
							"homeTeam": 0,
							"awayTeam": 0
						},
						"extraTime": {
							"homeTeam": null,
							"awayTeam": null
						},
						"penalties": {
							"homeTeam": null,
							"awayTeam": null
						}
					},
					"homeTeam": {
						"id": 8166,
						"name": "HB"
					},
					"awayTeam": {
						"id": 8912,
						"name": "Inter Club d'Escaldes"
					}
				}
			]
		}`)
	})

	season := Season{
		ID:              734,
		StartDate:       "2021-06-26",
		EndDate:         "2022-05-22",
		CurrentMatchday: 1,
	}

	match := Match{
		ID:          328846,
		Season:      &season,
		UtcDate:     "2021-06-22T18:00:00Z",
		Status:      string(StatusFinished),
		Stage:       "PRELIMINARY_ROUND",
		LastUpdated: "2021-09-04T16:20:05Z",
		Score: &Score{
			Winner:   "AWAY_TEAM",
			Duration: "REGULAR",
			FullTime: Time{
				HomeTeam: 0,
				AwayTeam: 1,
			},
			HalfTime: Time{
				HomeTeam: 0,
				AwayTeam: 0,
			},
			ExtraTime: Time{},
			Penalties: Time{},
		},
		HomeTeam: &Team{
			ID:   8166,
			Name: "HB",
		},
		AwayTeam: &Team{
			ID:   8912,
			Name: "Inter Club d'Escaldes",
		},
	}

	expected := &TeamMatches{
		Count:   1,
		Matches: []Match{match},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	list, err := client.Teams.Matches(ctx, "18", nil)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}
