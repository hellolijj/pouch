package client

import (
	"context"
	"fmt"
)

// ContainerDiff returns the file diffs of containers.
func (client *APIClient) ContainerDiff(ctx context.Context, name string) ([]string, error) {
	resp, err := client.get(ctx, "/containers/"+name+"changes", nil, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)

	diffs := []string{}
	err = decodeBody(&diffs, resp.Body)
	ensureCloseReader(resp)

	return diffs, err
}
