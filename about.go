package main

var aboutCommand = &Command {
        Name: "about",
        Implementation: func(cmd *Command, args []string) {
            printAbout(LongAbout)
        },
        Description: "Shows general information about this program",
        Help: `
    Example:

    $ deeq about
`,
    }
