package main

import "fmt"

var versionCommand = &Command {
        Name: "version",
        Implementation: func(cmd *Command, args []string) {
            fmt.Println(appVersion)
            cmd.GetApplication().SetSilence(true)
        },
        Description: "Prints the version of this program",
        Help: `
    Example:

    $ deeq version

`,
    }
