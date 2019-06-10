package client

import (
	"context"
)

// ContainerDiff returns the file diffs of containers.
func (client *APIClient) ContainerDiff(ctx context.Context, name string) ([]string, error) {
	resp, err := client.get(ctx, "/containers/"+name+"changes", nil, nil)

	if err != nil {
		return nil, err
	}

	diffs := []string{}
	err = decodeBody(&diffs, resp.Body)
	ensureCloseReader(resp)

	return diffs, err
}
