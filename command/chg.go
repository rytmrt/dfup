package command

import (
	"strings"
)

type ChgCommand struct {
	Meta
}

func (c *ChgCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ChgCommand) Synopsis() string {
	return "Change server infomation"
}

func (c *ChgCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
