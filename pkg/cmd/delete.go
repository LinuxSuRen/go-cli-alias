package cmd

import (
	"context"
	"github.com/linuxsuren/go-cli-alias/pkg"
	"github.com/spf13/cobra"
)

func NewDeleteCommand(ctx context.Context) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:  "delete",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			mgr := ctx.Value(pkg.AliasKey).(pkg.AliasManager)
			name := args[0]
			err = mgr.Delete(name)
			return
		},
	}
	return
}
