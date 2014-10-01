package main

import (
    prompt "github.com/bithavoc/goprompt"
    "fmt"
)

var forgotCommand = &Command {
        Name: "forgot",
        Implementation: func(cmd *Command, args []string) {
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "You are about to change your password using a password recovery code",
                        Fields: []*prompt.Field {
                            {
                                Name: "reset-code",
                                Title: "Reset Code",
                                Instructions: `Enter the reset code you receive in your email.
                                    If you haven't received one yet, make sure you use the 'recover' command to get one.
                                    `,
                                Shorthand: "rc",
                            },
                            {
                                Name: "password",
                                Title: "Password",
                                Instructions: "Enter the password for your account",
                                Shorthand: "p",
                            },
                            {
                                Name: "password_confirmation",
                                Title: "Password Confirmation",
                                Instructions: "Re-enter your password",
                                Shorthand: "pc",
                            },
                        },
                    },
                },
            }
            result := prompt.Process(args)
            form := result.Children["form.0"]
            code, password, password_confirmation := form.Children["reset-code"], form.Children["password"], form.Children["password_confirmation"]
            if password.Value != password_confirmation.Value {
                panic(fmt.Errorf("Password and Password confirmation do not match"))
            }
            app := cmd.GetDeeqApplication()
            if err := app.GetIdClient().Forgot(code.Value, password.Value); err != nil {
                panic(err)
            }
            fmt.Printf(`
    Password successfully changed.
`)
        },
        Description: "Change your password using a reset code",
        Help: `
    Changes your password using a reset code.

    Examples:
    
        $ deeq forgot
        > Reset Code: <emailed-code>
        > Password: <your_new_password>
        > Confirm Password: <your_new_password>

    or you could pass the values using arguments

        $ deeq forgot --reset-code=your_email@gmail.com --password=your_password --password_confirmation=your_password

    See the 'deeq signin' command for more information.
`,
    }
