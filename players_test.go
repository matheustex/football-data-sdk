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
