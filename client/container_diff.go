package client

import (
	"context"
	"net/url"
)

// ContainerDiff returns the file diffs of containers.
func (client *APIClient) ContainerDiff(ctx context.Context, name string) ([]string, error) {
	q := url.Values{}
	q.Set("name", name)

	resp, err := client.get(ctx, "/containers/changes", q, nil)
	if err != nil {
		return nil, err
	}

	diffs := []string{}
	err = decodeBody(&diffs, resp.Body)
	ensureCloseReader(resp)

	return diffs, err
}
