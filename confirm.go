package main

import (
    prompt "github.com/bithavoc/goprompt"
    "fmt"
)

var confirmCommand = &Command {
        Name: "confirm",
        Implementation: func(cmd *Command, args []string) {
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "Please enter the code you were given via email",
                        Fields: []*prompt.Field {
                            {
                                Name: "code",
                                Title: "Code",
                                Instructions: "Please enter your confirmation code",
                                Shorthand: "c",
                            },
                        },
                    },
                },
            }
            result := prompt.Process(args)
            form := result.Children["form.0"]
            code := form.Children["code"]

            app := cmd.GetDeeqApplication()
            _, err := app.GetIdClient().Confirm(code.Value)
            if err != nil {
                panic(err)
            }
            fmt.Printf(`
    Account successfully confirmed.

    You can log-in now.

    See 'deeq login' for more information.
`)
        },
        Description: "Confirms email and activates account",
        Help: `
    Once you sign-up for your Free account, you can use this command to confirm
    your account with the code that was sent to your email.

    Examples:
    
        $ deeq confirm
        > Code: 3b7f87c8f8a

    or you could pass the values using arguments:

        $ deeq confirm --code 3b7f87c8f8a

    also using shorthands:

        $ deeq confirm -c 3b7f87c8f8a

    After you confirm your account you will be able to log-in. Use the 'signin' command
    to start using Deeq.

    See the 'deeq signin' command for more information.
`,
    }
