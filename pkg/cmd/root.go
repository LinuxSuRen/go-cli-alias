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
	for _, v := range mgr.List() {
		//fmt.Println("register", k, v)
		root.AddCommand(&cobra.Command{
			Use:    v.Name,
			Hidden: true,
			RunE: func(cmd *cobra.Command, args []string) (err error) {
				rootName := root.Use

				var rootBinary string
				if rootBinary, err = exec.LookPath(rootName); err == nil {
					cmdArray := []string{rootBinary}
					cmdArray = append(cmdArray, strings.Split(v.Command, " ")...)
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
		Use:   "alias",
		Short: "Make your work more efficent by formula some wonderful command alias",
	}

	cmd.AddCommand(NewListCommand(ctx),
		NewSetCommand(ctx),
		NewDeleteCommand(ctx),
		NewInitCommand(ctx))
	return
}
