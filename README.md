dup -- diff-uploader [![Build Status](https://travis-ci.org/rytmrt/dup.svg)](https://travis-ci.org/rytmrt/dup)
====

## Description

Do updating by difference in the choose version and the current version.


## Usage


```
Usage: dup <command> [<args>]

  The command are as follows:
  deploy [-nD] [-r choose_rev]
    Do updating by difference in the choose revision and the current revision.

    The options are as follows:
    -n,--dry-run
      To order check contents of update.
    -D,--delete
      Has been deleted files so remove files on server.
    -r,--revision
      Specify the revision to deploy.

  init [-N server_name]
    Create skeleton for deplpyment.

    The options are as follows:
    -N,--server-name
      server name for deployment.  The default is empty.

  server-list
    Show server list.

  server-add <server_name> [-h hostname] [-p port] [-l login_name] [-i identity_file]
    Add server infomation.

    The arguments are as follows:
    server_name
      when you use of dup,  use this.

    The options are as follows:
    -h hostname, --host hostname
      Host name for server.  "user@hostname" is unsupported format.  The default is 'localhost'.

    -p port
      Port to connect to on the remote host.  The default port is "22".

    -l login_name
      Specify the user to login as the remote machine.

    -i identity_file
      Selects a file from which the identity (private key) for public key authentication is read.  The default is password authentication.

  server-chg <server_name> [-h hostname] [-p port] [-l login_name] [-i identity_file]
    Change server infomation.

    The options and the arguments see server-add command.

  server-del <server_name>
    Change server infomation.

    The options and the arguments see server-add command.

```

