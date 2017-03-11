package gogiphy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client - Giphy go client
type Client struct {
	Key        string
	HTTPS      bool
	Host       string
	Version    string
	HTTPClient *http.Client
}

// NewClient creates a new Giphy client with the specified key
// If no http.Client passed in, a new http.Client with default timeout of
// 30 seconds is created
func NewClient(key string, httpClient *http.Client) *Client {
	client := &Client{
		Key:     key,
		HTTPS:   false,
		Host:    "api.giphy.com",
		Version: "v1",
	}
	if httpClient != nil {
		client.HTTPClient = httpClient
	} else {
		client.HTTPClient = &http.Client{
			Timeout: time.Second * 30,
		}
	}
	return client
}

// Search for gifs. See https://github.com/Giphy/GiphyAPI#search-endpoint
func (client *Client) Search(query string, params *SearchParams) (*ListResponse, error) {
	scheme := "http"
	if client.HTTPS {
		scheme = "https"
	}
	searchURL := fmt.Sprintf("%s://%s/%s/gifs/search?api_key=%s&q=%s", scheme, client.Host, client.Version, client.Key, url.QueryEscape(query))
	if params != nil {
		segments := []string{searchURL}
		if params.Limit != 0 {
			segments = append(segments, fmt.Sprintf("&limit=%d", params.Limit))
		}
		if params.Offset != 0 {
			segments = append(segments, fmt.Sprintf("&offset=%d", params.Offset))
		}
		if params.Rating != "" {
			segments = append(segments, fmt.Sprintf("&rating=%s", url.QueryEscape(params.Rating)))
		}
		if params.Lang != "" {
			segments = append(segments, fmt.Sprintf("&lang=%s", url.QueryEscape(params.Lang)))
		}
		searchURL = strings.Join(segments, "")
	}
	data := &ListResponse{}
	err := get(client, searchURL, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func get(client *Client, url string, data interface{}) error {
	response, err := client.HTTPClient.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
