package main

import (
    prompt "github.com/bithavoc/goprompt"
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
                                DefaultValue: "true",
                            },
                        },
                    },
                },
            }
            result := prompt.Process()
            form := result.Children["form.0"]
            email, password, remember := form.Children["email"], form.Children["password"], form.Children["remember"]



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
