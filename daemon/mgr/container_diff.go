package mgr

import (
	"context"
	"fmt"
)

// List returns the container's list.
func (mgr *ContainerManager) List(ctx context.Context, name string) ([]string, error) {
	container, err := mgr.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	fmt.Println(container)

	return nil, nil
}
