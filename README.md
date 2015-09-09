dup -- diff-uploader
===

 [![Build Status](https://travis-ci.org/rytmrt/dup.svg)](https://travis-ci.org/rytmrt/dup)

Do updating by difference in the choose version and the current version

## Usage

### deploy
Do updating by difference in the choose revision and the current revision.

```bash
 $ dup deploy [-nD] [-r choose_rev]
```

The options are as follows

```
-n
  Dry run.  To order check contents of update.
-D
  Has been deleted files so remove files on server.
-r
  Specify the revision to deploy.
```

### init
Create skeleton for deplpyment.

```bash
 $ dup init [-N server_name]
```

The options are as follows

```
-N
server name for deployment.  The default is empty.
```

### list
Show server list.

```bash
 $ dup list
```

### add
Add server infomation.

```bash
 $ dup add <server_name> [-h hostname] [-p port] [-l login_name] [-i identity_file] [-P]
```

The arguments are as follows:

```
server_name
  when you use of dup,  use this.
```

The options are as follows

```
-h hostname
  Host name for server.  "user@hostname" is unsupported format.  The default is 'localhost'.

-p port
  Port to connect to on the remote host.  The default port is "22".

-l login_name
  Specify the user to login as the remote machine.

-i identity_file
  Selects a file from which the identity (private key) for public key authentication is read.  The default is password authentication.

-P
  Using password authentication.
```

### chg
Change server infomation.

```bash
$ dup chg <server_name> [-h hostname] [-p port] [-l login_name] [-i identity_file]
```

The options and the arguments see server-add command.

### del
Change server infomation.

```bash
$ dup del <server_name>
```

The options and the arguments see server-add command.

## Install

To install, use `go get`:

```bash
$ go get -d github.com/rytmrt/dup
```

## Contribution

1. Fork ([https://github.com/rytmrt/dup/fork](https://github.com/rytmrt/dup/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[rytmrt](https://github.com/rytmrt)
