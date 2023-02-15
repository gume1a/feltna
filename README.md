# Utility Tool For Gume1a
#### Installation

A valid [installation of Go](https://go.dev/doc/install) is required. Running the following command will result in allowance of running the commands from the integrated terminal.

```
go install github.com/gume1a/feltna@latest
```

This installs the latest version of the `feltna` cmd tool from the master branch.

#### Usage
Current usage is as displayed here. All of the stated command might have some optional flags that can be shown with running `feltna <command> --help` or shorthand `-h`.

```
Feltna is a cmd tool for Gume1a organization.

Usage:
  feltna [flags]
  feltna [command]

Available Commands:
  clone       command to clone all gume1a repositories
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for feltna

Use "feltna [command] --help" for more information about a command.
```

#### Developement
For development please refer to [cobra](https://github.com/spf13/cobra) package. It's used for command line tools.


