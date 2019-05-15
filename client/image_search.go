package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/alibaba/pouch/apis/types"
)

// ImageSearch requests daemon to search an image from registry.
func (client *APIClient) ImageSearch(ctx context.Context, term string, register string) ([]types.SearchResultItem, error) {
	var results []types.SearchResultItem

	q := url.Values{}
	q.Set("term", term)
	if len(register) > 0 {
		q.Set("registry", register)
	}

	// todo: add some auth info
	headers := map[string][]string{}

	resp, err := client.post(ctx, "/images/search", q, nil, headers)

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&results)
	return results, err
}
