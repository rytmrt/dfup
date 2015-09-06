dup - diff-uploader [![Build Status](https://travis-ci.org/rytmrt/dup.svg)](https://travis-ci.org/rytmrt/dup)
====

## Description

Do updating by difference in the choose version and the current version.


## Usage

```
Usage: dup [-nD] [tag] server_name


DESCRIPTIONS

    The options are as follows:
    -n,--dry-run
        To order check contents of update.
    -D,--delete
        Has been deleted files so remove files on server.
    tag
        contents in server for diff.
    server_name
        server name. need pre-registered server name.
```


```
Usage: dup-server [-a server_name hostname [-p port] [-l login_name] [-i identity_file]]
                  [-c server_name [hostname] [-p port] [-l login_name] [-i identity_file]]
                  [-D server_name]


DESCRIPTIONS

    The options are as follows:
    -a,--append
        Appending server info.

    -c,--change
        Changing server info of server_name.

    -D,--delete
        Deleting server info of server_name.

    server_name
        when you use of dup,  use this.

    hostname
        Host name for server.  "user@hostname" is unsupported format.

    -p port
        Port to connect to on the remote host.  The default port is "22".

    -l login_name
        Specify the user to login as the remote machine.

    -i identity_file
        Selects a file from which the identity (private key) for public key authentication is read.  The default is password authentication.
```

