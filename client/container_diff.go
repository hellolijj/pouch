package client

import (
	"context"
	"fmt"
)

// ContainerDiff returns the file diffs of containers.
func (client *APIClient) ContainerDiff(ctx context.Context, name string) ([]string, error) {
	fmt.Println("in ")
	resp, err := client.get(ctx, "/containers/"+name+"changes", nil, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(resp)

	diffs := []string{}
	err = decodeBody(&diffs, resp.Body)
	ensureCloseReader(resp)

	return diffs, err
}
