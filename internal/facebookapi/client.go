package facebookapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Me struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type Error struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FbtraceID string `json:"fbtrace_id"`
	} `json:"error"`
}

func GetMe(accessToken string) (*Me, error) {
	uri, err := url.Parse("https://graph.facebook.com/me")
	if err != nil {
		return nil, fmt.Errorf("parse url: %v", err)
	}

	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("parse query: %v", err)
	}
	query.Add("access_token", accessToken)
	uri.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		o := &Me{}
		if err = json.NewDecoder(res.Body).Decode(o); err != nil {
			return nil, err
		}
		return o, nil
	}

	e := &Error{}
	if err = json.NewDecoder(res.Body).Decode(e); err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("facebookapi.GetMe: %#v", e)
}
