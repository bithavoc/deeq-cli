package main

import (
    prompt "github.com/bithavoc/goprompt"
    id "github.com/bithavoc/id-go-client"
    "fmt"
)

var loginCommand = &Command {
        Name: "login",
        Implementation: func(cmd *Command, args []string) {
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "Enter your Login Credentials",
                        Fields: []*prompt.Field {
                            {
                                Name: "email",
                                Title: "Email",
                                Instructions: "Enter your bithavoc's email",
                            },
                            {
                                Name: "password",
                                Title: "Password",
                                Instructions: "Enter your bithavoc's password",
                            },
                            {
                                Name: "remember",
                                Title: "Remember",
                                Instructions: "Attach this session to your OS account?",
                                DefaultValue: "yes",
                            },
                        },
                    },
                },
            }
            result := prompt.Process()
            form := result.Children["form.0"]
            email, password, remember := form.Children["email"], form.Children["password"], form.Children["remember"]

            app := cmd.GetDeeqApplication()

            // login with user and password
            authCode, err := app.GetIdClient().LogIn(id.Credentials{
                email.Value,
                password.Value,
            })
            if err != nil {
                panic(err)
            }
            user, err := app.GetIdClient().Negotiate(authCode)
            if err != nil {
                panic(err)
            }
            err = app.SetCurrentUser(user, remember.Value == "yes")
            if err != nil {
                panic(err)
            }
            fmt.Printf("Welcome %s :)\n", app.GetCurrentUser().Info.Fullname)
        },
        Description: "Log-in using your Bithavoc's credentials",
        Help: `
    login with the given user and email

    Examples:
    
    $ deeq login
    > Email: your_email@gmail.com
    > Password: your_password

    $ deeq login --email=your_email@gmail.com --password=your_password

`,
    }
