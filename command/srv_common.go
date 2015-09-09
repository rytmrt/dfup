package command

import (
	"flag"
)

type ServerOptions struct {
		serverName string
		hostName string
		port int
		password bool
		loginName string
		identityFile string
}

func parseServerOption(args []string) (options ServerOptions) {

	// Define option flag parse
	flags := flag.NewFlagSet("srv_option", flag.ContinueOnError)
	flags.Usage = func() {}

	flags.StringVar(&options.hostName, "H", "localhost", "Host name")
	flags.IntVar(&options.port, "p", 22, "Port")
	flags.StringVar(&options.loginName, "l", "root", "Login name")
	flags.StringVar(&options.identityFile, "i", "~/.ssh/id_rsa", "Login name")
	flags.BoolVar(&options.password, "P", false, "Password authentication")

	var help bool
	flags.BoolVar(&help, "h", false, "Show help")


	// Parse commandline flag
	if err := flags.Parse(args[0:]); err != nil {
		var r ServerOptions
		r.serverName = "ERROR"
		return r
	}
	var arguments []string
	for 0 < flags.NArg() {
		arguments = append(arguments, flags.Arg(0))
		flags.Parse(flags.Args()[1:])
	}

	if help {
		options.serverName = "SHOW_HELP"
	} else {
		options.serverName = arguments[0]
	}

	return
}
