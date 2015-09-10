package command

import (
	"testing"
)

func TestParseServerOption_implement(t *testing.T) {

	testCase := [][]string{
		{"-h"},
		{"server_name_00", "-H", "host.name"},
		{"server_name_01", "-p", "23"},
		{"server_name_02", "-l", "admin"},
		{"server_name_03", "-i", "~/identity_file"},
		{"server_name_04", "-P"},
		{"server_name_05", "-H", "host.name", "-p", "23"},
		{"server_name_06", "-H", "host.name", "-i", "~/identity_file"},
		{"server_name_07", "-H", "host.name", "-l", "admin"},
		{"server_name_08", "-p", "23", "-l", "admin"},
		{"server_name_09", "-p", "23", "-i", "~/identity_file"},
		{"server_name_10", "-l", "admin", "-i", "~/identity_file"},
		{"server_name_11", "-H", "host.name", "-p", "23", "-l", "admin"},
		{"server_name_12", "-H", "host.name", "-p", "23", "-i", "~/identity_file"},
		{"server_name_13", "-H", "host.name", "-l", "admin", "-i", "~/identity_file"},
		{"server_name_14", "-p", "23", "-l", "admin", "-i", "~/identity_file"},
		{"server_name_15", "-H", "host.name", "-p", "23", "-l", "admin", "-i", "~/identity_file"},
		{"server_name_16", "-H", "host.name", "-p", "23", "-P"},
		{"server_name_17", "-H", "host.name", "-i", "~/identity_file", "-P"},
		{"server_name_18", "-H", "host.name", "-l", "admin", "-P"},
		{"server_name_19", "-p", "23", "-l", "admin", "-P"},
		{"server_name_20", "-p", "23", "-i", "~/identity_file", "-P"},
		{"server_name_21", "-l", "admin", "-i", "~/identity_file", "-P"},
		{"server_name_22", "-H", "host.name", "-p", "23", "-l", "admin", "-P"},
		{"server_name_23", "-H", "host.name", "-p", "23", "-i", "~/identity_file", "-P"},
		{"server_name_24", "-H", "host.name", "-l", "admin", "-i", "~/identity_file", "-P"},
		{"server_name_25", "-p", "23", "-l", "admin", "-i", "~/identity_file", "-P"},
		{"server_name_26", "-H", "host.name", "-p", "23", "-l", "admin", "-i", "~/identity_file", "-P"},
	}

	
	testResult := []ServerOptions{
		ServerOptions{ serverName: "SHOW_HELP",      hostName: "localhost", port: 22, loginName: "root",  identityFile: "~/.ssh/id_rsa",   password: false},
		ServerOptions{ serverName: "server_name_00", hostName: "host.name", port: 22, loginName: "root",  identityFile: "~/.ssh/id_rsa",   password: false},
		ServerOptions{ serverName: "server_name_01", hostName: "localhost", port: 23, loginName: "root",  identityFile: "~/.ssh/id_rsa",   password: false},
		ServerOptions{ serverName: "server_name_02", hostName: "localhost", port: 22, loginName: "admin", identityFile: "~/.ssh/id_rsa",   password: false},
		ServerOptions{ serverName: "server_name_03", hostName: "localhost", port: 22, loginName: "root",  identityFile: "~/identity_file", password: false},
		ServerOptions{ serverName: "server_name_04", hostName: "localhost", port: 22, loginName: "root",  identityFile: "~/.ssh/id_rsa",   password: true },
		ServerOptions{ serverName: "server_name_05", hostName: "host.name", port: 23, loginName: "root",  identityFile: "~/.ssh/id_rsa",   password: false},
		ServerOptions{ serverName: "server_name_06", hostName: "host.name", port: 22, loginName: "root",  identityFile: "~/identity_file", password: false},
		ServerOptions{ serverName: "server_name_07", hostName: "host.name", port: 22, loginName: "admin", identityFile: "~/.ssh/id_rsa",   password: false},
		ServerOptions{ serverName: "server_name_08", hostName: "localhost", port: 23, loginName: "admin", identityFile: "~/.ssh/id_rsa",   password: false},
		ServerOptions{ serverName: "server_name_09", hostName: "localhost", port: 23, loginName: "root",  identityFile: "~/identity_file", password: false},
		ServerOptions{ serverName: "server_name_10", hostName: "localhost", port: 22, loginName: "admin", identityFile: "~/identity_file", password: false},
		ServerOptions{ serverName: "server_name_11", hostName: "host.name", port: 23, loginName: "admin", identityFile: "~/.ssh/id_rsa",   password: false},
		ServerOptions{ serverName: "server_name_12", hostName: "host.name", port: 23, loginName: "root",  identityFile: "~/identity_file", password: false},
		ServerOptions{ serverName: "server_name_13", hostName: "host.name", port: 22, loginName: "admin", identityFile: "~/identity_file", password: false},
		ServerOptions{ serverName: "server_name_14", hostName: "localhost", port: 23, loginName: "admin", identityFile: "~/identity_file", password: false},
		ServerOptions{ serverName: "server_name_15", hostName: "host.name", port: 23, loginName: "admin", identityFile: "~/identity_file", password: false},
		ServerOptions{ serverName: "server_name_16", hostName: "host.name", port: 23, loginName: "root",  identityFile: "~/.ssh/id_rsa",   password: true },
		ServerOptions{ serverName: "server_name_17", hostName: "host.name", port: 22, loginName: "root",  identityFile: "~/identity_file", password: true },
		ServerOptions{ serverName: "server_name_18", hostName: "host.name", port: 22, loginName: "admin", identityFile: "~/.ssh/id_rsa",   password: true },
		ServerOptions{ serverName: "server_name_19", hostName: "localhost", port: 23, loginName: "admin", identityFile: "~/.ssh/id_rsa",   password: true },
		ServerOptions{ serverName: "server_name_20", hostName: "localhost", port: 23, loginName: "root",  identityFile: "~/identity_file", password: true },
		ServerOptions{ serverName: "server_name_21", hostName: "localhost", port: 22, loginName: "admin", identityFile: "~/identity_file", password: true },
		ServerOptions{ serverName: "server_name_22", hostName: "host.name", port: 23, loginName: "admin", identityFile: "~/.ssh/id_rsa",   password: true },
		ServerOptions{ serverName: "server_name_23", hostName: "host.name", port: 23, loginName: "root",  identityFile: "~/identity_file", password: true },
		ServerOptions{ serverName: "server_name_24", hostName: "host.name", port: 22, loginName: "admin", identityFile: "~/identity_file", password: true },
		ServerOptions{ serverName: "server_name_25", hostName: "localhost", port: 23, loginName: "admin", identityFile: "~/identity_file", password: true },
		ServerOptions{ serverName: "server_name_26", hostName: "host.name", port: 23, loginName: "admin", identityFile: "~/identity_file", password: true },
	}

	for i, v := range testCase {
		svrOptions := parseServerOption(v)
		if svrOptions.serverName != testResult[i].serverName {
			t.Errorf("missing serverName.  wrong: %v.  test case: %#v", svrOptions.serverName, v)
		}
		if svrOptions.hostName != testResult[i].hostName {
			t.Errorf("missing hostName.  wrong: %v.  test case: %#v", svrOptions.hostName, v)
		}
		if svrOptions.port != testResult[i].port {
			t.Errorf("missing port.  wrong: %v.  test case: %#v", svrOptions.port, v)
		}
		if svrOptions.loginName != testResult[i].loginName {
			t.Errorf("missing loginName.  wrong: %v.  test case: %#v", svrOptions.loginName, v)
		}
		if svrOptions.identityFile != testResult[i].identityFile {
			t.Errorf("missing identityFile.  wrong: %v.  test case: %#v", svrOptions.identityFile, v)
		}
		if svrOptions.password != testResult[i].password {
			t.Errorf("missing password.  wrong: %v.  test case: %#v", svrOptions.password, v)
		}
	}
}
