package tsumby

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Image struct {
	Data []byte
	Type string
}

func (api *API) Create(ctx context.Context, p Params) (*Image, error) {
	path := Generate(p, NewHMACSigner(sha256.New, 40, api.APISecret))

	// do HTTP request
	uri := fmt.Sprintf("%s", path)

	var respBody []byte
	var err error

	resp, respErr := api.request(ctx, http.MethodGet, uri, nil, nil)
	// short circuit processing on context timeouts
	if respErr != nil && errors.Is(respErr, context.DeadlineExceeded) {
		return nil, respErr
	}

	if respErr != nil || resp.StatusCode >= 500 {
		if respErr == nil {
			return nil, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
		} else {
			return nil, fmt.Errorf("HTTP request failed: %w", respErr)
		}
	} else {
		respBody, err = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		if err != nil {
			return nil, fmt.Errorf("could not read response body: %w", err)
		}
	}

	// return response body and content type from header resp
	return &Image{
		Data: respBody,
		Type: resp.Header.Get("Content-Type"),
	}, nil

}
