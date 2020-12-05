package cmd

import (
	"context"
	"github.com/linuxsuren/go-cli-alias/pkg"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func RegisterAliasCommands(ctx context.Context, root *cobra.Command) {
	//fmt.Println(ctx.Value(pkg.AliasKey))
	mgr := ctx.Value(pkg.AliasKey).(pkg.AliasManager)
	//fmt.Println(len(mgr.List()), "==")
	for k, v := range mgr.List() {
		//fmt.Println("register", k, v)
		root.AddCommand(&cobra.Command{
			Use:    k,
			Hidden: true,
			RunE: func(cmd *cobra.Command, args []string) (err error) {
				rootName := root.Use

				var rootBinary string
				if rootBinary, err = exec.LookPath(rootName); err == nil {
					cmdArray := []string{rootBinary}
					cmdArray = append(cmdArray, strings.Split(v, " ")...)
					cmdArray = append(cmdArray, args...)
					err = syscall.Exec(rootBinary, cmdArray, os.Environ())
				}
				return
			},
		})
	}
}

func NewRootCommand(ctx context.Context) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "alias",
	}

	cmd.AddCommand(NewListCommand(ctx),
		NewSetCommand(ctx),
		NewDeleteCommand(ctx))
	return
}