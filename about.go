package main

var aboutCommand = &Command {
        Name: "about",
        Implementation: func(cmd *Command, args []string) {
            printAbout(LongAbout)
        },
        Description: "General information about this command",
        Help: `
    Example:

    $ deeq about
`,
    }
