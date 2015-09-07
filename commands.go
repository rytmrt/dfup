package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/rytmrt/dup/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "deploy",
		Usage:  "Deploy to diff",
		Action: command.CmdDeploy,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "init",
		Usage:  "Create skeleton for deplpyment",
		Action: command.CmdInit,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "list",
		Usage:  "Show server list",
		Action: command.CmdList,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "add",
		Usage:  "Add server infomation",
		Action: command.CmdAdd,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "chg",
		Usage:  "Change server infomation",
		Action: command.CmdChg,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "del",
		Usage:  "Change server infomation",
		Action: command.CmdDel,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
