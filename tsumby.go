package tsumby

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

var (
	Version string = "dev"
)

// API holds the configuration for the API client.
type API struct {
	APISecret  string
	UserAgent  string
	BaseURL    string
	headers    http.Header
	httpClient *http.Client
}

func New(secret string) *API {
	return &API{
		BaseURL:    fmt.Sprintf("%s://%s%s", defaultScheme, defaultHostname, defaultBasePath),
		APISecret:  secret,
		UserAgent:  userAgent + "/" + Version,
		headers:    make(http.Header),
		httpClient: http.DefaultClient,
	}
}

func (api *API) request(ctx context.Context, method, uri string, reqBody io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, api.BaseURL+uri, reqBody)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, api.headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	if api.UserAgent != "" {
		req.Header.Set("User-Agent", api.UserAgent)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	return resp, nil
}

// copyHeader copies all headers for `source` and sets them on `target`.
// based on https://godoc.org/github.com/golang/gddo/httputil/header#Copy
func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}
