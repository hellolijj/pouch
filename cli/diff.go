package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// diffDescription is used to describe diff command in detail and auto generate command doc.
var diffDescription = "Inspect changes to files or directories on a container's filesystem"

// DiffCommand use to implement 'diff' command, it is used to look for changes to files
type DiffCommand struct {
	baseCommand
}

// Init initialize diff command.
func (diff *DiffCommand) Init(c *Cli) {
	diff.cli = c
	diff.cmd = &cobra.Command{
		Use:   "diff [OPTIONS] CONTAINER",
		Short: "look for changes to files",
		Long:  diffDescription,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return diff.runDiff(args)
		},
		Example: diffExample(),
	}
}

// runDiff is the entry of DiffCommand command.
func (diff *DiffCommand) runDiff(args []string) error {
	containerName := args[0]

	ctx := context.Background()
	apiClient := diff.cli.Client()

	diffs, err := apiClient.ContainerDiff(ctx, containerName)
	if err != nil {
		return err
	}
	fmt.Println(diffs)
	return nil
}

// diffExample shows examples in diff command, and is used in auto-generated cli docs.
func diffExample() string {
	return `$ pouch diff e747c540815f
C /usr
C /usr/local
A /usr/local/aegis
A /usr/local/aegis/aegis_client
A /usr/local/aegis/aegis_client/aegis_10_65
A /usr/local/aegis/aegis_client/aegis_10_65/data
A /usr/local/aegis/aegis_client/aegis_10_65/data/data.5
A /usr/local/aegis/aegis_client/aegis_10_65/data/data.6
A /usr/local/aegis/aegis_client/aegis_10_65/data/data.7
A /usr/local/aegis/aegis_client/aegis_10_65/data/data.1
A /usr/local/aegis/aegis_client/aegis_10_65/data/data.2
A /usr/local/aegis/aegis_client/aegis_10_65/data/data.3
A /usr/local/aegis/aegis_client/aegis_10_65/data/data.4
C /run
A /run/nginx.pid
C /var
C /var/cache
C /var/cache/nginx
A /var/cache/nginx/client_temp
A /var/cache/nginx/fastcgi_temp
A /var/cache/nginx/proxy_temp
A /var/cache/nginx/scgi_temp
A /var/cache/nginx/uwsgi_temp`
}
