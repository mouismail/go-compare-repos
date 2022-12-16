package github

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/go-logr/logr"
	"github.com/google/go-github/v47/github"
	"github.com/gregjones/httpcache"
	"golang.org/x/oauth2"

	"migrator/build"
)

// Config contains configuration for Github client
type Config struct {
	EnterpriseURL     string `split_words:"true"`
	AppID             int64  `split_words:"true"`
	AppInstallationID int64  `split_words:"true"`
	AppPrivateKey     string `split_words:"true"`
	Token             string
	URL               string `split_words:"true"`
	UploadURL         string `split_words:"true"`
	BasicauthUsername string `split_words:"true"`
	BasicauthPassword string `split_words:"true"`
	RunnerGitHubURL   string `split_words:"true"`

	Log *logr.Logger
}

// Client wraps GitHub client with some additional
type Client struct {
	*github.Client
	regTokens     map[string]*github.RegistrationToken
	mu            sync.Mutex
	GithubBaseURL string
	IsEnterprise  bool
}

type BasicAuthTransport struct {
	Username string
	Password string
}

func (p BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(p.Username, p.Password)
	return http.DefaultTransport.RoundTrip(req)
}

// NewClient creates a GitHub Client
func (c *Config) NewClient() (*Client, error) {
	var transport http.RoundTripper

	if len(c.BasicauthUsername) > 0 && len(c.BasicauthPassword) > 0 {
		transport = BasicAuthTransport{Username: c.BasicauthUsername, Password: c.BasicauthPassword}
	} else if len(c.Token) > 0 {
		transport = oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: c.Token})).Transport
	} else {
		var tr *ghinstallation.Transport

		if _, err := os.Stat(c.AppPrivateKey); err == nil {
			tr, err = ghinstallation.NewKeyFromFile(http.DefaultTransport, c.AppID, c.AppInstallationID, c.AppPrivateKey)
			if err != nil {
				return nil, fmt.Errorf("authentication failed: using private key at %s: %v", c.AppPrivateKey, err)
			}
		} else {
			tr, err = ghinstallation.New(http.DefaultTransport, c.AppID, c.AppInstallationID, []byte(c.AppPrivateKey))
			if err != nil {
				return nil, fmt.Errorf("authentication failed: using private key of size %d (%s...): %v", len(c.AppPrivateKey), strings.Split(c.AppPrivateKey, "\n")[0], err)
			}
		}

		if len(c.EnterpriseURL) > 0 {
			githubAPIURL, err := getEnterpriseApiUrl(c.EnterpriseURL)
			if err != nil {
				return nil, fmt.Errorf("enterprise url incorrect: %v", err)
			}
			tr.BaseURL = githubAPIURL
		}
		transport = tr
	}

	cached := httpcache.NewTransport(httpcache.NewMemoryCache())
	cached.Transport = transport
	httpClient := &http.Client{}

	var client *github.Client
	var githubBaseURL string
	var isEnterprise bool
	if len(c.EnterpriseURL) > 0 {
		var err error
		isEnterprise = true
		client, err = github.NewEnterpriseClient(c.EnterpriseURL, c.EnterpriseURL, httpClient)
		if err != nil {
			return nil, fmt.Errorf("enterprise client creation failed: %v", err)
		}
		githubBaseURL = fmt.Sprintf("%s://%s%s", client.BaseURL.Scheme, client.BaseURL.Host, strings.TrimSuffix(client.BaseURL.Path, "api/v3/"))
	} else {
		client = github.NewClient(httpClient)
		githubBaseURL = "https://github.com/"

		if len(c.URL) > 0 {
			baseUrl, err := url.Parse(c.URL)
			if err != nil {
				return nil, fmt.Errorf("github client creation failed: %v", err)
			}
			if !strings.HasSuffix(baseUrl.Path, "/") {
				baseUrl.Path += "/"
			}
			client.BaseURL = baseUrl
		}

		if len(c.UploadURL) > 0 {
			uploadUrl, err := url.Parse(c.UploadURL)
			if err != nil {
				return nil, fmt.Errorf("github client creation failed: %v", err)
			}
			if !strings.HasSuffix(uploadUrl.Path, "/") {
				uploadUrl.Path += "/"
			}
			client.UploadURL = uploadUrl
		}

		if len(c.RunnerGitHubURL) > 0 {
			githubBaseURL = c.RunnerGitHubURL
			if !strings.HasSuffix(githubBaseURL, "/") {
				githubBaseURL += "/"
			}
		}
	}
	client.UserAgent = "go-action-runner/" + build.Version
	return &Client{
		Client:        client,
		regTokens:     map[string]*github.RegistrationToken{},
		mu:            sync.Mutex{},
		GithubBaseURL: githubBaseURL,
		IsEnterprise:  isEnterprise,
	}, nil
}

// cleanup removes expired registration tokens.
func (c *Client) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, rt := range c.regTokens {
		if rt.GetExpiresAt().Before(time.Now()) {
			delete(c.regTokens, key)
		}
	}
}

func (c *Client) createRegistrationToken(ctx context.Context, enterprise, org, repo string) (*github.RegistrationToken, *github.Response, error) {
	if len(repo) > 0 {
		return c.Client.Actions.CreateRegistrationToken(ctx, org, repo)
	}
	if len(org) > 0 {
		return c.Client.Actions.CreateOrganizationRegistrationToken(ctx, org)
	}
	return c.Client.Enterprise.CreateRegistrationToken(ctx, enterprise)
}

func getEnterpriseApiUrl(baseURL string) (string, error) {
	baseEndpoint, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/api/v3/") &&
		!strings.HasPrefix(baseEndpoint.Host, "api.") &&
		!strings.Contains(baseEndpoint.Host, ".api.") {
		baseEndpoint.Path += "api/v3/"
	}

	// Trim trailing slash, otherwise there's double slash added to token endpoint
	return fmt.Sprintf("%s://%s%s", baseEndpoint.Scheme, baseEndpoint.Host, strings.TrimSuffix(baseEndpoint.Path, "/")), nil
}

type RunnerNotFound struct {
	runnerName string
}

func (e *RunnerNotFound) Error() string {
	return fmt.Sprintf("runner %q not found", e.runnerName)
}

type RunnerOffline struct {
	runnerName string
}

func (e *RunnerOffline) Error() string {
	return fmt.Sprintf("runner %q offline", e.runnerName)
}
