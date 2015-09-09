package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDelCommand_implement(t *testing.T) {
	var _ cli.Command = &DelCommand{}
}
