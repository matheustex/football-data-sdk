package football

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	APIURL = "https://api.football-data.org/v2"
)

type Client struct {
	client  *http.Client // HTTP client used to communicate with the API.
	BaseURL *url.URL

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Football API.
	Areas        *AreasService
	Competitions *CompetitionService
	Matches      *MatchService
	Players      *PlayersService
	Teams        *TeamsService
}

type service struct {
	client *Client
}

// NewClient returns a new GitHub API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(APIURL)

	c := &Client{client: httpClient, BaseURL: baseURL}

	c.common.client = c

	c.Areas = (*AreasService)(&c.common)
	c.Competitions = (*CompetitionService)(&c.common)
	c.Matches = (*MatchService)(&c.common)
	c.Players = (*PlayersService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)

	return c
}

// Get performs a GET against the api
func (c *Client) Get(path string, params interface{}, v interface{}) ([]byte, error) {
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

	headers.Set("X-Auth-Token", "")

	return *headers
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
