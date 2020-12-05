package cmd

import (
	"context"
	"github.com/linuxsuren/go-cli-alias/pkg"
	"github.com/spf13/cobra"
)

func NewListCommand(ctx context.Context) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "list",
		Short: "list all alias command lines",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			mgr := ctx.Value(pkg.AliasKey).(pkg.AliasManager)
			var list map[string]string
			list = mgr.List()
			cmd.Println(list)
			return
		},
	}
	return
}
