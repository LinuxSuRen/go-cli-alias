package cmd

import (
	"context"
	ext "github.com/linuxsuren/cobra-extension/pkg"
	"github.com/linuxsuren/go-cli-alias/pkg"
	"github.com/spf13/cobra"
)

func NewListCommand(ctx context.Context) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "list",
		Short: "list all alias command lines",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			mgr := ctx.Value(pkg.AliasKey).(pkg.AliasManager)
			var list []pkg.Alias
			list = mgr.List()

			out := ext.OutputOption{
				Writer:  cmd.OutOrStdout(),
				Columns: "Name,Command",
			}
			err = out.OutputV2(list)
			return
		},
	}
	return
}
