package tinderapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
	SelfID     string
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

	x, err := client.GetProfile()
	if err != nil {
		return nil, err
	}

	if x.Meta.Status != http.StatusOK {
		return nil, fmt.Errorf("Cannot initialize the Tinder REST API client: %v", x)
	}

	client.Status = fmt.Sprintf("Success!\nUser name: %s\nUser ID: %s",
		x.Data.Account.Username, x.Data.User.ID)

	client.SelfID = x.Data.User.ID

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
	include := []string{
		"account",
		"boost",
		"contact_cards",
		"email_settings",
		"instagram",
		"likes",
		"notifications",
		"plus_control",
		"products",
		"purchase",
		"readreceipts",
		"swipenote",
		"spotify",
		"super_likes",
		"tinder_u",
		"travel",
		"tutorials",
		"user",
	}
	query.Add("include", strings.Join(include, ","))
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

// GetMatches returns matches
func (c *Client) GetMatches() (*Page, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/matches", c.BaseURL))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("parse query: %v", err)
	}
	query.Add("count", "60")
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

// GetMessages returns messages by given matchID match ID
func (c *Client) GetMessages(matchID string) (*Page, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/matches/%s/messages", matchID,
		c.BaseURL))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("parse query: %v", err)
	}
	query.Add("count", "100")
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

// SendMessage sends a message with given text text to matchID match ID
func (c *Client) SendMessage(matchID string, text string) (*Message, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/user/matches/%s", matchID,
		c.BaseURL))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	message := &Message{Message: text}
	b, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", uri.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res := Message{}
	if err = c.sendRequest(req, &res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return &res, nil
}

// Like likes a user by given user ID id
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
	uri, err := url.Parse(fmt.Sprintf("%s/pass/%s", c.BaseURL, id))
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
