package main

import (
  "github.com/mitchellh/cli"
  "github.com/rytmrt/dup/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
  return map[string]cli.CommandFactory{
    "deploy": func() (cli.Command, error) {
      return &command.DeployCommand{
        Meta: *meta,
      }, nil
    },
    "init": func() (cli.Command, error) {
      return &command.InitCommand{
        Meta: *meta,
      }, nil
    },
    "list": func() (cli.Command, error) {
      return &command.ListCommand{
        Meta: *meta,
      }, nil
    },
    "add": func() (cli.Command, error) {
      return &command.AddCommand{
        Meta: *meta,
      }, nil
    },
    "chg": func() (cli.Command, error) {
      return &command.ChgCommand{
        Meta: *meta,
      }, nil
    },
    "del": func() (cli.Command, error) {
      return &command.DelCommand{
        Meta: *meta,
      }, nil
    },

    "version": func() (cli.Command, error) {
      return &command.VersionCommand{
        Meta:     *meta,
        Version:  Version,
        Revision: GitCommit,
        Name: Name,
      }, nil
    },
  }
}
