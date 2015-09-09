package command

import (
	"strings"
)

type DelCommand struct {
	Meta
}

func (c *DelCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *DelCommand) Synopsis() string {
	return "Delete server infomation"
}

func (c *DelCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
