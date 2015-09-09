package command

import (
	"strings"
	"fmt"
)

type AddCommand struct {
	Meta
}

func (c *AddCommand) Run(args []string) int {

	fmt.Printf("%#v\n\n", args)

	var (
		options ServerOptions
		argsErr bool = false
	)

	if len(args) < 0 {
		options = parseServerOption(args)
	} else {
		argsErr = true
	}


	if argsErr || options.serverName == "ERROR" || options.serverName == "" || options.serverName == "SHOW_HELP" {
		c.Ui.Output(c.Help())
	} else {
		fmt.Printf("serverName   : %v\n", options.serverName)
		fmt.Printf("hostName     : %v\n", options.hostName)
		fmt.Printf("port         : %v\n", options.port)
		fmt.Printf("loginName    : %v\n", options.loginName)
		fmt.Printf("identityFile : %v\n", options.identityFile)
		fmt.Printf("password     : %v\n", options.password)
	}

	return 0
}

func (c *AddCommand) Synopsis() string {
	return "Add server infomation"
}

func (c *AddCommand) Help() string {
	helpText := `
Add server infomation.

[Usage]
	dup add <server_name> [-h hostname] [-p port] [-l login_name] [-i identity_file] [-P]

The arguments are as follows:
	server_name
	  when you use of dup,  use this.

The options are as follows:
	-H hostname
	  Host name for server.  "user@hostname" is unsupported format.
	  The default is 'localhost'.

	-p port
	  Port to connect to on the remote host.
	  The default is 22.

	-l login_name
	  Specify the user to login as the remote machine.
	  The default is 'root'.

	-i identity_file
	  Selects a file from which the identity (private key) for public
	  key authentication is read.  The default is "~/.ssh/id_rsa".

	-P
	  Using password authentication.
`
	return strings.TrimSpace(helpText)
}
