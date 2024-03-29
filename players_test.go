package football

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerService_Find(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/players/1", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"id": 1,
			"name": "Illarramendi",
			"firstName": "Asier",
			"lastName": null,
			"dateOfBirth": "1990-03-08",
			"countryOfBirth": "Spain",
			"nationality": "Spain",
			"position": "Midfielder",
			"shirtNumber": null,
			"lastUpdated": "2020-09-07T21:26:05Z"
		}`)
	})

	expected := &Player{
		ID:             1,
		Name:           "Illarramendi",
		FirstName:      "Asier",
		DateOfBirth:    "1990-03-08",
		CountryOfBirth: "Spain",
		Nationality:    "Spain",
		Position:       "Midfielder",
		LastUpdated:    "2020-09-07T21:26:05Z",
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	player, err := client.Players.Find(ctx, "1")

	assert.Nil(t, err)
	assert.Equal(t, expected, player)
}

func TestPlayerService_FindWithError(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/players/55555", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{
			"message": "The resource you are looking for does not exist.",
			"error": 404
		}`)
	})

	expected := errors.New("404 Not Found")

	ctx := context.Background()
	client := NewClient(httpClient)
	_, err := client.Players.Find(ctx, "55555")

	assert.NotNil(t, err)

	if !ErrorContains(err, expected.Error()) {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestPlayerService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/players/18/matches", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"count": 1,
			"player": {
				"id": 18,
				"name": "Joaquín",
				"firstName": "Joaquín",
				"dateOfBirth": "1981-07-21",
				"countryOfBirth": "Spain",
				"nationality": "Spain",
				"position": "Midfielder",
				"lastUpdated": "2020-11-26T02:18:33Z"
			},
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

	player := &Player{
		ID:             18,
		Name:           "Joaquín",
		FirstName:      "Joaquín",
		DateOfBirth:    "1981-07-21",
		CountryOfBirth: "Spain",
		Nationality:    "Spain",
		Position:       "Midfielder",
		LastUpdated:    "2020-11-26T02:18:33Z",
	}

	expected := &PlayerMatches{
		Count:   1,
		Player:  *player,
		Matches: []Match{match},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	list, err := client.Players.Matches(ctx, "18", nil)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}
