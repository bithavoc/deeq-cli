package main

import (
    "fmt"
)

var whoamiCommand = &Command {
        Name: "whoami",
        Implementation: func(cmd *Command, args []string) {
            app := cmd.GetDeeqApplication()
            user := app.GetCurrentUser()
            fmt.Println("Name:", user.Info.Fullname)
            fmt.Println("Email:", user.Info.Email)
            fmt.Println("Token:", user.Token.Code)
        },
        Description: "Shows information about the current user",
        Help: `

    If you already logged in using deeq login, you can use this command to know more about you :)

    Examples:
    
    $ deeq whoami

`,
    }
