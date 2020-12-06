[![](https://goreportcard.com/badge/linuxsuren/go-cli-alias)](https://goreportcard.com/report/linuxsuren/go-cli-alias)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/linuxsuren/go-cli-alias)
[![Contributors](https://img.shields.io/github/contributors/linuxsuren/go-cli-alias.svg)](https://github.com/linuxsuren/go-cli-alias/graphs/contributors)
[![GitHub release](https://img.shields.io/github/release/linuxsuren/go-cli-alias.svg?label=release)](https://github.com/linuxsuren/go-cli-alias/releases/latest)
![GitHub All Releases](https://img.shields.io/github/downloads/linuxsuren/go-cli-alias/total)

# Go CLI Alias

Adding a command alias feature for your CLI.

# Get started

`go get github.com/linuxsuren/go-cli-alias`

Put the following code lines:

```
var ctx context.Context
if defMgr, err := alias.GetDefaultAliasMgr(); err == nil {
    ctx = context.WithValue(context.Background(), alias.AliasKey, defMgr)

    rootCmd.AddCommand(cmd.NewRootCommand(ctx))

    cmd.RegisterAliasCommands(ctx, rootCmd)
} else {
    fmt.Println(fmt.Errorf("cannot get default alias manager, error: %v", err))
}
```
