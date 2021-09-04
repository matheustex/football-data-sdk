package football

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchService_Find(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/matches/1", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"head2head": {
				"numberOfMatches": 1,
				"totalGoals": 3,
				"homeTeam": {
					"wins": 0,
					"draws": 0,
					"losses": 1
				},
				"awayTeam": {
					"wins": 1,
					"draws": 0,
					"losses": 0
				}
			},			
			"match": {
				"id": 204950,
				"competition": {
					"id": 2013,
					"name": "Série A"
				},
				"season": {
					"id": 15,
					"startDate": "2018-04-14",
					"endDate": "2018-12-02",
					"currentMatchday": 25,
					"availableStages": [
						"REGULAR_SEASON"
					]
				},
				"utcDate": "2018-08-13T23:00:00Z",
				"status": "FINISHED",
				"minute": null,
				"attendance": 14640,
				"venue": "Estadio Jornalista Mário Filho",
				"matchday": 18,
				"stage": "REGULAR_SEASON",
				"group": "Regular Season",
				"lastUpdated": "2018-08-19T20:01:26Z",
				"homeTeam": {
					"id": 1765,
					"name": "Fluminense FC",
					"coach": {
						"id": 73229,
						"name": "Marcelo Oliveira",
						"countryOfBirth": "Brazil",
						"nationality": "Brazil"
					},
					"captain": {
						"id": 1068,
						"name": "Gum",
						"shirtNumber": 3
					},
					"lineup": [
						{
							"id": 1083,
							"name": "Jádson",
							"position": "Midfielder",
							"shirtNumber": 16
						}
					],
					"bench": [
						{
							"id": 1083,
							"name": "Jádson",
							"position": "Midfielder",
							"shirtNumber": 16
						}
					]
				},
				"awayTeam": {
					"id": 1765,
					"name": "Fluminense FC",
					"coach": {
						"id": 73229,
						"name": "Marcelo Oliveira",
						"countryOfBirth": "Brazil",
						"nationality": "Brazil"
					},
					"captain": {
						"id": 1068,
						"name": "Gum",
						"shirtNumber": 3
					},
					"lineup": [
						{
							"id": 1083,
							"name": "Jádson",
							"position": "Midfielder",
							"shirtNumber": 16
						}
					],
					"bench": [
						{
							"id": 1083,
							"name": "Jádson",
							"position": "Midfielder",
							"shirtNumber": 16
						}
					]
				},
				"score": {
					"winner": "AWAY_TEAM",
					"duration": "REGULAR",
					"fullTime": {
						"homeTeam": 0,
						"awayTeam": 3
					},
					"halfTime": {
						"homeTeam": 0,
						"awayTeam": 3
					}					
				},
				"goals": [
					{
						"minute": 23,						
						"type": "REGULAR",
						"team": {
							"id": 6684,
							"name": "SC Internacional"
						},
						"scorer": {
							"id": 1588,
							"name": "Nicolás López"
						},
						"assist": {
							"id": 1575,
							"name": "Rodrigo Dourado"
						}
					}
				],
				"bookings": [],
				"substitutions": [],
				"referees": []
			}
		}`)
	})

	player := Player{
		ID:          1083,
		Name:        "Jádson",
		Position:    "Midfielder",
		ShirtNumber: 16,
	}

	team := Team{
		ID:   1765,
		Name: "Fluminense FC",
		Coach: &Coach{
			ID:             73229,
			Name:           "Marcelo Oliveira",
			CountryOfBirth: "Brazil",
			Nationality:    "Brazil",
		},
		Captain: &Player{
			ID:          1068,
			Name:        "Gum",
			ShirtNumber: 3,
		},
		Lineup: &[]Player{player},
		Bench:  &[]Player{player},
	}

	score := Score{
		Winner:   "AWAY_TEAM",
		Duration: "REGULAR",
		FullTime: Time{
			HomeTeam: 0,
			AwayTeam: 3,
		},
		HalfTime: Time{
			HomeTeam: 0,
			AwayTeam: 3,
		},
	}

	match := Match{
		ID: 204950,
		Competition: &Competition{
			ID:   2013,
			Name: "Série A",
		},
		Season: &Season{
			ID:              15,
			StartDate:       "2018-04-14",
			EndDate:         "2018-12-02",
			CurrentMatchday: 25,
			AvailableStages: []string{"REGULAR_SEASON"},
		},
		UtcDate:     "2018-08-13T23:00:00Z",
		Status:      string(StatusFinished),
		Attendance:  14640,
		Venue:       "Estadio Jornalista Mário Filho",
		Matchday:    18,
		Stage:       "REGULAR_SEASON",
		Group:       "Regular Season",
		LastUpdated: "2018-08-19T20:01:26Z",
		HomeTeam:    &team,
		AwayTeam:    &team,
		Score:       &score,
		Goals: []Goals{
			{
				Minute: 23,
				Type:   "REGULAR",
				Team: Team{
					ID:   6684,
					Name: "SC Internacional",
				},
				Scorer: Player{
					ID:   1588,
					Name: "Nicolás López",
				},
				Assist: Player{
					ID:   1575,
					Name: "Rodrigo Dourado",
				},
			},
		},
		Bookings:      []Bookings{},
		Substitutions: []Substitutions{},
		Referees:      []Referees{},
	}

	head2head := Head2Head{
		NumberOfMatches: 1,
		TotalGoals:      3,
		HomeTeam: TeamStats{
			Wins:   0,
			Draws:  0,
			Losses: 1,
		},
		AwayTeam: TeamStats{
			Wins:   1,
			Draws:  0,
			Losses: 0,
		},
	}

	expected := &MatchResponse{
		Head2Head: head2head,
		Match:     match,
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	matchResponse, err := client.Matches.Find(ctx, "1")

	assert.Nil(t, err)
	assert.Equal(t, expected, matchResponse)
}

func TestMatchService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/matches", func(w http.ResponseWriter, r *http.Request) {
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

	expected := &MatchesCompetition{
		Count:   1,
		Matches: []Match{match},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	list, err := client.Matches.List(ctx, nil)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}
