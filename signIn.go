package main

import (
    prompt "github.com/bithavoc/goprompt"
    id "github.com/bithavoc/id-go-client"
    "fmt"
)

var signInCommand = &Command {
        Name: "signin",
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
                                Shorthand: "e",
                            },
                            {
                                Name: "password",
                                Title: "Password",
                                Instructions: "Enter your bithavoc's password",
                                Shorthand: "p",
                            },
                            {
                                Name: "remember",
                                Title: "Remember",
                                Instructions: "Attach this session to your OS account?",
                                DefaultValue: "yes",
                                Shorthand: "r",
                            },
                        },
                    },
                },
            }
            result := prompt.Process(args)
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
        Description: "Signs-in user with email and password",
        Help: `
    Signs-in user with email and password.

    Examples:
    
        $ deeq signin
        > Email: your_email@gmail.com
        > Password: your_password

    or you could pass the values using arguments

        $ deeq signin --email=your_email@gmail.com --password=your_password

    Use the 'deeq signup' command to create your account if you don't have one yet.
`,
    }
