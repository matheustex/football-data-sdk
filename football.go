package football

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/google/go-querystring/query"
)

const (
	APIURL = "https://api.football-data.org/v2"
)

type Client struct {
	client  *http.Client // HTTP client used to communicate with the API.
	BaseURL *url.URL
	common  service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Football API.
	Areas        *AreaService
	Competitions *CompetitionService
	Matches      *MatchService
	Players      *PlayerService
	Teams        *TeamService
}

type service struct {
	client *Client
}

// NewClient returns a new GitHub API client. If a nil httpClient is
// provided, a new http.Client will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(APIURL)

	c := &Client{client: httpClient, BaseURL: baseURL}

	c.common.client = c

	c.Areas = (*AreaService)(&c.common)
	c.Competitions = (*CompetitionService)(&c.common)
	c.Matches = (*MatchService)(&c.common)
	c.Players = (*PlayerService)(&c.common)
	c.Teams = (*TeamService)(&c.common)

	return c
}

// Get performs a GET against the api
func (c *Client) Get(path string, params interface{}, v interface{}) ([]byte, error) {
	if len(os.Getenv("FOOTBALL_API_TOKEN")) == 0 {
		return nil, errors.New("You need to export the FOOTBALL_API_TOKEN")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", APIURL, path), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {
		query, _ := query.Values(params)

		req.URL.RawQuery = query.Encode()
	}

	req.Header = c.GetHeaders()

	res, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		err := errors.New(res.Status)
		return nil, err
	}

	if response != nil {
		json.Unmarshal(response, &v)
	}

	return response, nil
}

func (client *Client) GetHeaders() http.Header {
	headers := &http.Header{}

	headers.Set("X-Auth-Token", os.Getenv("FOOTBALL_API_TOKEN"))

	return *headers
}
