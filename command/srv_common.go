package command

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os/user"
	"strings"
)

const (
	LocationPath = "~/.dup/server_options/"
	Extenstion   = ".toml"
)

type ServerOptions struct {
	serverName   string
	hostName     string
	port         int
	password     bool
	loginName    string
	identityFile string
}

func ParseServerOption(args []string) (options ServerOptions) {

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

func ConvPath(srcPath string) string {
	usr, _ := user.Current()
	r := strings.Replace(srcPath, "~", usr.HomeDir, 1)
	return r
}

func SaveServerOptionsFile(serverOptions ServerOptions) (e error) {
	return
}

func LoadingServerOptionsFile(serverName string) (serverOptions ServerOptions, e error) {
	return
}

func Toml2ServerOptions(str string) (serverOptions ServerOptions, e error) {
	return
}

func ServerOptions2Toml(serverOptions *ServerOptions) (toml string, e error) {
	return
}

func ServerUpdate(server ServerOptions, modList []string, addList []string, delList []string) (e error) {
	// cennect server -----
	key, err := GetKeyFile(server.identityFile)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: server.serverName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}

	fmt.Printf("%#v", *config)

	// upload to server -----
	// delete to server -----
	return
}

func GetKeyFile(path string) (key ssh.Signer, err error) {
	file := ConvPath(path)
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		return
	}
	return
}
