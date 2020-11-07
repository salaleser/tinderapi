package tinderapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/salaleser/tinderapi/cmd/structs"
	"github.com/salaleser/tinderapi/internal/pkg/facebookapi"
)

const (
	// BaseURL is the Tinder API URL
	BaseURL = "https://api.gotinder.com"
)

// Client represents a Tinder REST API client
type Client struct {
	BaseURL    string
	Status     string
	SelfID     string
	HTTPClient *http.Client
	token      string
}

// NewClient returns a Tinder REST API client.
func NewClient() *Client {
	return &Client{
		BaseURL: BaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// LoginFacebook authorizes a client via Facebook.
func (c *Client) LoginFacebook(facebookToken string) error {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/auth/login/facebook", c.BaseURL))
	if err != nil {
		return fmt.Errorf("parse url: %v", err)
	}

	type Credentials struct {
		Token      string `json:"token"`
		FacebookID string `json:"facebook_id"`
	}

	me, err := facebookapi.GetMe(facebookToken)
	if err != nil {
		return fmt.Errorf("login: %v", err)
	}

	credentials := &Credentials{Token: facebookToken, FacebookID: me.ID}
	b, err := json.Marshal(credentials)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, uri.String(),
		bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res := &structs.Page{}
	if err = c.sendRequest(req, res); err != nil {
		return fmt.Errorf("http request: %v", err)
	}

	c.token = res.Data.APIToken

	profile, err := c.GetProfile()
	if err != nil {
		return err
	}

	if profile.Meta.Status != http.StatusOK {
		return fmt.Errorf("unable to authenticate a Tinder REST API client: %v",
			profile)
	}

	c.Status = fmt.Sprintf("*Success!*\n*User name:* %s\n*User ID:* %s",
		profile.Data.Account.Username, profile.Data.User.ID)

	c.SelfID = profile.Data.User.ID

	return nil
}

// Login authorizes a client via Tinder token.
func (c *Client) Login(token string) error {
	c.token = token

	profile, err := c.GetProfile()
	if err != nil {
		return err
	}

	if profile.Meta.Status != http.StatusOK {
		return fmt.Errorf(
			"unable to authenticate a Tinder REST API client: %v", profile)
	}

	c.Status = fmt.Sprintf("*Success!*\n*User name:* %s\n*User ID:* %s",
		profile.Data.Account.Username, profile.Data.User.ID)

	c.SelfID = profile.Data.User.ID

	return nil
}

// GetUser returns Tinder user by given ID
func (c *Client) GetUser(id string) (*structs.UserV1, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/user/%s", c.BaseURL, id))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &structs.UserV1{}
	if err = c.sendRequest(req, res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return res, nil
}

// GetProfile returns profile information
func (c *Client) GetProfile() (*structs.Page, error) {
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

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &structs.Page{}
	if err = c.sendRequest(req, res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return res, nil
}

// SetProfile sets profile information.
func (c *Client) SetProfile(user structs.User) (*structs.Page, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/profile", c.BaseURL))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	b, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, uri.String(),
		bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res := &structs.Page{}
	if err = c.sendRequest(req, res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return res, nil
}

// GetRecommendations returns recommendations
func (c *Client) GetRecommendations() (*structs.Page, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/recs/core", c.BaseURL))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &structs.Page{}
	if err = c.sendRequest(req, res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return res, nil
}

// GetMatches returns matches
func (c *Client) GetMatches(count int, message int) (*structs.Page, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/matches", c.BaseURL))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("parse query: %v", err)
	}
	query.Add("count", strconv.Itoa(count))
	query.Add("message", strconv.Itoa(message))
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &structs.Page{}
	if err = c.sendRequest(req, res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return res, nil
}

// GetMessages returns messages by given matchID match ID
func (c *Client) GetMessages(matchID string, count int) (*structs.Page, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/v2/matches/%s/messages", c.BaseURL,
		matchID))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("parse query: %v", err)
	}
	query.Add("count", strconv.Itoa(count))
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &structs.Page{}
	if err = c.sendRequest(req, res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return res, nil
}

// SendMessage sends a message with given text text to matchID match ID
func (c *Client) SendMessage(matchID string, text string) (*structs.Message, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/user/matches/%s", c.BaseURL,
		matchID))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	message := structs.Message{Message: text}
	b, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, uri.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res := &structs.Message{}
	if err = c.sendRequest(req, res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return res, nil
}

// Like likes a user by given user ID id
func (c *Client) Like(id string) (*structs.Like, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/like/%s", c.BaseURL, id))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &structs.Like{}
	if err = c.sendRequest(req, res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return res, nil
}

// Pass passes user by given ID
func (c *Client) Pass(id string) (*structs.Pass, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/pass/%s", c.BaseURL, id))
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &structs.Pass{}
	if err = c.sendRequest(req, res); err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	return res, nil
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
