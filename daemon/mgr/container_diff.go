package mgr

import (
	"context"
	"fmt"
)

// Diff returns the container's file's differences.
func (mgr *ContainerManager) Diff(ctx context.Context, name string) ([]string, error) {
	container, err := mgr.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	fmt.Println(container)

	return nil, nil
}
