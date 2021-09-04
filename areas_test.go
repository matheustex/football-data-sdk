package football

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaService_Find(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/areas/1", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"id": 2000,
			"name": "Afghanistan",
			"countryCode": "AFG",
			"ensignUrl": null,
			"parentAreaId": 2014,
			"parentArea": "Asia",
			"childAreas": []
		}`)
	})

	expected := Area{
		ID:           2000,
		Name:         "Afghanistan",
		CountryCode:  "AFG",
		ParentAreaID: 2014,
		ParentArea:   "Asia",
		ChildAreas:   &[]Area{},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	area, err := client.Areas.Find(ctx, "1")

	assert.Nil(t, err)
	assert.Equal(t, expected, *area)
}

func TestAreaService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/areas", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"count": 273,
			"areas": [
				{
					"id": 2000,
					"name": "Afghanistan",
					"countryCode": "AFG",
					"ensignUrl": null,
					"parentAreaId": 2014,
					"parentArea": "Asia"
				}
			]
		}`)
	})

	expected := &AreaList{
		Count: 273,
		Areas: []Area{
			{
				ID:           2000,
				Name:         "Afghanistan",
				CountryCode:  "AFG",
				ParentAreaID: 2014,
				ParentArea:   "Asia",
			},
		},
	}

	ctx := context.Background()
	client := NewClient(httpClient)
	list, err := client.Areas.List(ctx)

	assert.Nil(t, err)
	assert.Equal(t, expected, list)
}
