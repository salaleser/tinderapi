package tinderapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	// BaseURL is the Tinder API URL
	BaseURL = "https://api.gotinder.com"
)

// Client represents the Tinder API client
type Client struct {
	BaseURL    string
	Status     string
	HTTPClient *http.Client
	token      string
	endpoint   string
}

// New initializes and returns the Tinder API client
func New(token string) (*Client, error) {
	client := &Client{
		BaseURL: BaseURL,
		token:   token,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	page, err := client.GetProfile()
	if err != nil {
		return nil, err
	}

	client.Status = fmt.Sprintf("Client successfully initialized for user *%s* (`%s`)",
		page.Data.Account.Username, page.Data.User.ID)

	return client, nil
}

// GetUser returns Tinder user by given ID
func (c *Client) GetUser(id string) (*UserV1, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/user/%s", c.BaseURL, id))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := UserV1{}
	if err = c.sendRequest(req, &res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return &res, nil
}

// GetProfile returns profile information
func (c *Client) GetProfile() (*Page, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/profile", c.BaseURL))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("parse query: %v", err)
	}
	query.Add("include", "account,user")
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := Page{}
	if err = c.sendRequest(req, &res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return &res, nil
}

// GetRecommendations returns recommendations
func (c *Client) GetRecommendations() (*Page, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/recs/core", c.BaseURL))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := Page{}
	if err = c.sendRequest(req, &res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return &res, nil
}

// Like likes user by given ID
func (c *Client) Like(id string) (*Like, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/like/%s", c.BaseURL, id))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := Like{}
	if err = c.sendRequest(req, &res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return &res, nil
}

// Pass passes user by given ID
func (c *Client) Pass(id string) (*Pass, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/like/%s", c.BaseURL, id))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := Pass{}
	if err = c.sendRequest(req, &res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return &res, nil
}

func (c *Client) sendRequest(req *http.Request, o interface{}) error {
	req.Header.Add("X-Auth-Token", c.token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&o); err != nil {
		return err
	}

	return nil
}
