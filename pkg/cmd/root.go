package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/linuxsuren/go-cli-alias/pkg"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func CreateDefaultCmd(target, alias string) *cobra.Command {
	return &cobra.Command{
		Use:  alias,
		RunE: DefaultRunE(target),
	}
}

func DefaultRunE(targetCLI string) func(cmd *cobra.Command, args []string) (err error) {
	return func(cmd *cobra.Command, args []string) (err error) {
		env := os.Environ()

		var targetArgs []string
		targetCLIArray := strings.Split(targetCLI, " ")
		if len(targetCLIArray) > 1 {
			targetCLI = targetCLIArray[0]
			targetArgs = targetCLIArray[1:]
			targetArgs = append(targetArgs, args...)
		} else {
			targetArgs = args
		}

		var gitBinary string
		if gitBinary, err = exec.LookPath(targetCLI); err == nil {
			err = syscall.Exec(gitBinary, append([]string{targetCLI}, targetArgs...), env)
		}
		return
	}
}

func RedirectToAlias(ctx context.Context, args []string) (redirect bool, aliasCmd []string) {
	if len(args) <= 0 {
		return
	}

	mgr := ctx.Value(pkg.AliasKey).(pkg.AliasManager)
	for _, v := range mgr.List() {
		if v.Name == args[0] {
			redirect = true
			aliasCmd = strings.Split(v.Command, " ")
			if len(args) > 1 {
				aliasCmd = append(aliasCmd, args[1:]...)
			}
			break
		}
	}
	return
}

func RegisterAliasCompletion(ctx context.Context, root *cobra.Command) {
	mgr := ctx.Value(pkg.AliasKey).(pkg.AliasManager)
	aliasNames := []string{}
	for _, v := range mgr.List() {
		aliasNames = append(aliasNames, v.Name)
	}
	root.ValidArgsFunction = func(cmd *cobra.Command, args []string, toComplete string) (i []string, directive cobra.ShellCompDirective) {
		return aliasNames, cobra.ShellCompDirectiveNoFileComp
	}
}

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
		Short: "Make your work more efficient by formula some wonderful command alias",
	}

	cmd.AddCommand(NewListCommand(ctx),
		NewSetCommand(ctx),
		NewDeleteCommand(ctx),
		NewInitCommand(ctx))
	return
}

func AddAliasCmd(cmd *cobra.Command, defaultAlias []pkg.Alias) {
	var ctx context.Context
	if defMgr, err := pkg.GetDefaultAliasMgrWithNameAndInitialData(cmd.Name(), defaultAlias); err == nil {
		ctx = context.WithValue(context.Background(), pkg.AliasKey, defMgr)

		cmd.AddCommand(NewRootCommand(ctx))
	} else {
		cmd.Println(fmt.Errorf("cannot get default alias manager, error: %v", err))
	}
}

func Execute(cmd *cobra.Command, target string, aliasList []pkg.Alias, preHook func([]string)) {
	ExecuteContext(cmd, context.TODO(), target, aliasList, preHook)
}

func ExecuteContext(cmd *cobra.Command, ctx context.Context, target string, aliasList []pkg.Alias, preHook func([]string)) {
	cmd.SilenceErrors = true
	var err error
	// this is very trick way, looking to improve it
	if len(strings.Split(target, " ")) > 1 {
		err = errors.New("unknown command")
	} else if ctx != nil {
		err = cmd.ExecuteContext(ctx)
	} else {
		err = cmd.Execute()
	}

	if err != nil {
		if strings.Contains(err.Error(), "unknown command") {
			args := os.Args[1:]
			var defMgr *pkg.DefaultAliasManager
			if defMgr, err = pkg.GetDefaultAliasMgrWithNameAndInitialData(cmd.Name(), aliasList); err == nil {
				ctx := context.WithValue(context.Background(), pkg.AliasKey, defMgr)
				var targetBinary string
				var targetArgs []string
				var targetCmd string
				env := os.Environ()

				targetCmdArray := strings.Split(target, " ")
				if len(targetCmdArray) > 1 {
					targetCmd = targetCmdArray[0]
					targetArgs = targetCmdArray[1:]
				} else {
					targetCmd = targetCmdArray[0]
				}

				if targetBinary, err = exec.LookPath(targetCmd); err != nil {
					panic(fmt.Sprintf("cannot find %s", target))
				}

				if ok, redirect := RedirectToAlias(ctx, args); ok {
					args = redirect
				}

				if preHook != nil {
					preHook(args)
				}

				targetArgs = append(targetCmdArray, args...)
				_ = syscall.Exec(targetBinary, targetArgs, env) // ignore the errors due to we've no power to deal with it
			} else {
				err = fmt.Errorf("cannot get default alias manager, error: %v", err)
			}
		} else {
			cmd.PrintErrln(err)
		}
	}
}
