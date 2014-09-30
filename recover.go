package main

import (
    prompt "github.com/bithavoc/goprompt"
    "fmt"
)

var recoverCommand = &Command {
        Name: "recover",
        Implementation: func(cmd *Command, args []string) {
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "Enter your email address to get instructions to reset your email",
                        Fields: []*prompt.Field {
                            {
                                Name: "email",
                                Title: "Email",
                                Instructions: "Please enter the email address of your account",
                                Shorthand: "e",
                            },
                        },
                    },
                },
            }
            result := prompt.Process(args)
            form := result.Children["form.0"]
            email := form.Children["email"]

            app := cmd.GetDeeqApplication()
            if err := app.GetIdClient().Recover(email.Value); err != nil {
                panic(err)
            }
            fmt.Printf(`
    Check your email inbox for further instructions.

    Then use the 'forgot' command to reset your password.
`)
        },
        Description: "Get instructions to reset your password",
        Help: `
    If you forgot your password and you do remember your email address,
    you can use this command to get instructions to reset your password.

    Examples:
    
        $ deeq recover
        > Email: john.doe@example.com

    or you could pass the values using arguments:

        $ deeq recover --email john.doe@example.com

    also using shorthands:

        $ deeq recover -e john.doe@example.com

    As part of the instructions in the email, there is a code you can use
    with the command 'forgot'.

    See the 'deeq forgot' command for more information.
`,
    }
