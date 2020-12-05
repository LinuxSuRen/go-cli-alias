package cmd

import (
	"context"
	"github.com/linuxsuren/go-cli-alias/pkg"
	"github.com/spf13/cobra"
)

func NewSetCommand(ctx context.Context) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:  "set",
		Args: cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			mgr := ctx.Value(pkg.AliasKey).(pkg.AliasManager)
			name, cli := args[0], args[1]
			err = mgr.Set(name, cli)
			return
		},
	}
	return
}
