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
		Name:   "server-list",
		Usage:  "Show server list",
		Action: command.CmdServer - List,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "server-add",
		Usage:  "Add server infomation",
		Action: command.CmdServer - Add,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "server-chg",
		Usage:  "Change server infomation",
		Action: command.CmdServer - Chg,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "server-del",
		Usage:  "Change server infomation",
		Action: command.CmdServer - Del,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
