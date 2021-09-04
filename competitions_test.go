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

func TestCompetitionService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/competitions", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"count": 1,
			"competitions": [
				{
					"id": 2006,
					"area": {
						"id": 2001,
						"name": "Africa",
						"countryCode": "AFR",
						"ensignUrl": null
					},
					"name": "WC Qualification",
					"code": null,
					"emblemUrl": null,
					"plan": "TIER_FOUR",
					"currentSeason": {
						"id": 555,
						"startDate": "2019-09-04",
						"endDate": "2021-11-16",
						"currentMatchday": 4,
						"winner": null
					},
					"numberOfAvailableSeasons": 2,
					"lastUpdated": "2018-06-04T23:54:04Z"
				}				
			]
		}`)
	})

	season := Season{
		ID:              555,
		StartDate:       "2019-09-04",
		EndDate:         "2021-11-16",
		CurrentMatchday: 4,
	}

	competition := &Competition{
		ID: 2006,
		Area: Area{
			ID:          2001,
			Name:        "Africa",
			CountryCode: "AFR",
		},
		Name:                     "WC Qualification",
		Plan:                     "TIER_FOUR",
		CurrentSeason:            season,
		NumberOfAvailableSeasons: 2,
		LastUpdated:              "2018-06-04T23:54:04Z",
	}

	filters := CompetitionFiltersOptions{
		Areas: "2001",
	}

	expected := &CompetitionList{
		Count:        1,
		Competitions: []Competition{*competition},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	list, err := client.Competitions.List(ctx, &filters)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}

func TestCompetitionService_Teams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/competitions/2001/teams", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"count": 1,			
			"competition": {
				"id": 2001,
				"area": {
					"id": 2077,
					"name": "Europe"
				},
				"name": "UEFA Champions League",
				"code": "CL",
				"plan": "TIER_ONE",
				"lastUpdated": "2021-06-26T13:37:26Z"
			},
			"season": {
				"id": 734,
				"startDate": "2021-06-26",
				"endDate": "2022-05-22",
				"currentMatchday": 1,
				"winner": null
			},
			"teams": [
				{
					"id": 4,
					"area": {
						"id": 2088,
						"name": "Germany"
					},
					"name": "Borussia Dortmund",
					"shortName": "Dortmund",
					"tla": "BVB",
					"crestUrl": "https://crests.football-data.org/4.svg",
					"address": "Rheinlanddamm 207-209 Dortmund 44137",
					"phone": "+49 (231) 90200",
					"website": "http://www.bvb.de",
					"email": "info@bvb.de",
					"founded": 1909,
					"clubColors": "Black / Yellow",
					"venue": "Signal Iduna Park",
					"lastUpdated": "2021-04-14T07:43:46Z"
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

	competition := Competition{
		ID: 2001,
		Area: Area{
			ID:   2077,
			Name: "Europe",
		},
		Name:        "UEFA Champions League",
		Code:        "CL",
		Plan:        "TIER_ONE",
		LastUpdated: "2021-06-26T13:37:26Z",
	}

	team := Team{
		ID: 4,
		Area: &Area{
			ID:   2088,
			Name: "Germany",
		},
		Name:        "Borussia Dortmund",
		ShortName:   "Dortmund",
		Tla:         "BVB",
		CrestURL:    "https://crests.football-data.org/4.svg",
		Address:     "Rheinlanddamm 207-209 Dortmund 44137",
		Phone:       "+49 (231) 90200",
		Website:     "http://www.bvb.de",
		Email:       "info@bvb.de",
		Founded:     1909,
		ClubColors:  "Black / Yellow",
		Venue:       "Signal Iduna Park",
		LastUpdated: "2021-04-14T07:43:46Z",
	}

	expected := &CompetitionTeams{
		Count:       1,
		Competition: competition,
		Season:      season,
		Teams:       []Team{team},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	list, err := client.Competitions.Teams(ctx, "2001", nil)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}

func TestCompetitionService_Standings(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/competitions/2001/standings", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"count": 1,			
			"competition": {
				"id": 2001,
				"area": {
					"id": 2077,
					"name": "Europe"
				},
				"name": "UEFA Champions League",
				"code": "CL",
				"plan": "TIER_ONE",
				"lastUpdated": "2021-06-26T13:37:26Z"
			},
			"season": {
				"id": 734,
				"startDate": "2021-06-26",
				"endDate": "2022-05-22",
				"currentMatchday": 1,
				"winner": null
			},
			"standings": [
				{
					"stage": "GROUP_STAGE",
					"type": "TOTAL",
					"group": "PRELIMINARY_ROUND",
					"table": []
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

	competition := Competition{
		ID: 2001,
		Area: Area{
			ID:   2077,
			Name: "Europe",
		},
		Name:        "UEFA Champions League",
		Code:        "CL",
		Plan:        "TIER_ONE",
		LastUpdated: "2021-06-26T13:37:26Z",
	}

	standing := Standing{
		Stage: "GROUP_STAGE",
		Type:  "TOTAL",
		Group: "PRELIMINARY_ROUND",
		Table: []Table{},
	}

	expected := &CompetitionStandings{
		Competition: competition,
		Season:      season,
		Standings:   []Standing{standing},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	list, err := client.Competitions.Standings(ctx, "2001", nil)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}

func TestCompetitionService_Matches(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/competitions/2001/matches", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"count": 1,			
			"competition": {
				"id": 2001,
				"area": {
					"id": 2077,
					"name": "Europe"
				},
				"name": "UEFA Champions League",
				"code": "CL",
				"plan": "TIER_ONE",
				"lastUpdated": "2021-06-26T13:37:26Z"
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

	competition := Competition{
		ID: 2001,
		Area: Area{
			ID:   2077,
			Name: "Europe",
		},
		Name:        "UEFA Champions League",
		Code:        "CL",
		Plan:        "TIER_ONE",
		LastUpdated: "2021-06-26T13:37:26Z",
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

	expected := &CompetitionMatches{
		Count:       1,
		Competition: competition,
		Matches:     []Match{match},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	list, err := client.Competitions.Matches(ctx, "2001", nil)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}

func TestCompetitionService_Scorers(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/competitions/2001/scorers", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"count": 1,			
			"competition": {
				"id": 2001,
				"area": {
					"id": 2077,
					"name": "Europe"
				},
				"name": "UEFA Champions League",
				"code": "CL",
				"plan": "TIER_ONE",
				"lastUpdated": "2021-06-26T13:37:26Z"
			},
			"season": {
				"id": 734,
				"startDate": "2021-06-26",
				"endDate": "2022-05-22",
				"currentMatchday": 1,
				"winner": null
			},
			"scorers": [
				{
					"player": {
						"id": 16596,
						"name": "Antonio-Mirko Čolak",
						"firstName": "Antonio-Mirko",
						"lastName": null,
						"dateOfBirth": "1993-09-17",
						"countryOfBirth": "Germany",
						"nationality": "Croatia",
						"position": "Attacker",
						"shirtNumber": 17,
						"lastUpdated": "2020-09-03T03:38:46Z"
					},
					"team": {
						"id": 749,
						"name": "Malmö FF"
					},
					"numberOfGoals": 5
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

	competition := Competition{
		ID: 2001,
		Area: Area{
			ID:   2077,
			Name: "Europe",
		},
		Name:        "UEFA Champions League",
		Code:        "CL",
		Plan:        "TIER_ONE",
		LastUpdated: "2021-06-26T13:37:26Z",
	}

	scorer := Scorer{
		Player: Player{
			ID:             16596,
			Name:           "Antonio-Mirko Čolak",
			FirstName:      "Antonio-Mirko",
			DateOfBirth:    "1993-09-17",
			CountryOfBirth: "Germany",
			Nationality:    "Croatia",
			Position:       "Attacker",
			ShirtNumber:    17,
			LastUpdated:    "2020-09-03T03:38:46Z",
		},
		Team: Team{
			ID:   749,
			Name: "Malmö FF",
		},
		NumberOfGoals: 5,
	}

	expected := &CompetitionScorers{
		Count:       1,
		Competition: competition,
		Season:      season,
		Scorers:     []Scorer{scorer},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	list, err := client.Competitions.Scorers(ctx, "2001", nil)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}
