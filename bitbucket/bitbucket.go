package bitbucket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CloneURL    string `json:"clone_url"`
	IsPrivate   bool   `json:"is_private"`
}

type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
	Username   string
	Password   string
}

func NewClient(baseURL string, username string, password string) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		BaseURL:    u,
		httpClient: http.DefaultClient,
		Username:   username,
		Password:   password,
	}

	return c, nil
}

func (c *Client) GetRepos() ([]*Repo, error) {
	// Construct the URL for the request.
	u := *c.BaseURL
	u.Path = "/api/2.0/repositories"

	// Create the request.
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Set the authentication headers.
	req.SetBasicAuth(c.Username, c.Password)

	// Make the request.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check the status code.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d: %s", resp.StatusCode, string(body))
	}

	// Parse the JSON response.
	var result struct {
		Values []*Repo `json:"values"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Values, nil
}
