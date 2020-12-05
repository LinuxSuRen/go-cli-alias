package main

import (
	"context"
	"fmt"
	"github.com/linuxsuren/go-cli-alias/pkg"
	"github.com/linuxsuren/go-cli-alias/pkg/cmd"
	"github.com/spf13/cobra"
	"log"
)

func NewRootCommand() (root *cobra.Command) {
	root = &cobra.Command{
		Use:   "ga",
		Short: "alias your command lines",
	}

	var ctx context.Context
	if defMgr, err := pkg.GetDefaultAliasMgr(); err == nil {
		ctx = context.WithValue(context.Background(), pkg.AliasKey, defMgr)

		root.AddCommand(cmd.NewRootCommand(ctx))

		cmd.RegisterAliasCommands(ctx, root)
	} else {
		log.Println(fmt.Errorf("cannot get default alias manager, error: %v", err))
	}
	return
}
