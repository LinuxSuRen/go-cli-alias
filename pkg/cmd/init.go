package cmd

import (
	"context"
	"github.com/linuxsuren/go-cli-alias/pkg"
	"github.com/spf13/cobra"
)

func NewInitCommand(ctx context.Context) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "init",
		Short: "init the pre-defined alia commands",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			mgr := ctx.Value(pkg.AliasKey).(pkg.AliasManager)
			err = mgr.Init()
			return
		},
	}
	return
}
