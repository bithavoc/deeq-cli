package main

import (
    prompt "github.com/bithavoc/goprompt"
    id "github.com/bithavoc/id-go-client"
    "fmt"
)

var logoutCommand = &Command {
        Name: "logout",
        RequiresUser: true,
        Implementation: func(cmd *Command, args []string) {
            app := cmd.GetDeeqApplication()
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "You are about to log out",
                        Fields: []*prompt.Field {
                            {
                                Name: "sure",
                                Title: "Do you really want to logout? (yes/no)",
                                Instructions: "Please tell me if you really want to log out, type either yes or no",
                                Shorthand: "s",
                            },
                        },
                    },
                },
            }
            result := prompt.Process(args)
            form := result.Children["form.0"]
            logoutAnswer := form.Children["sure"]
            if logoutAnswer.Value == "yes" {
                err := app.SetCurrentUser(id.User{}, true)
                if err != nil {
                    panic(err)
                }
                fmt.Println("Ok, Bye :(")
            } else {
                fmt.Printf("Thank you %s, good to have you here :)\n", app.GetCurrentUser().Info.Fullname)
            }
        },
        Description: "Logs out the currently logged-in user",
        Help: `

        If you already logged in using deeq login, you can use this command to forget yourself

    Examples:
    
    $ deeq logout
    > Ok, Bye :(

    Of course, you don't want to do this... you want to use Deeq!... you want to use Deeq! *fades away*
`,
    }
