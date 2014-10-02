package main

import (
    prompt "github.com/bithavoc/goprompt"
    id "github.com/bithavoc/id-go-client"
    "fmt"
)

var signOutCommand = &Command {
        Name: "signout",
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
        Description: "Signs-out the current user",
        Help: `

        Use this command to sign-out.

    Examples:
    
        $ deeq signout

    or to log-out without confirmation:

        $ deeq signou --sure=yes

`,
    }
