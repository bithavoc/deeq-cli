package main

import (
    prompt "github.com/bithavoc/goprompt"
    id "github.com/bithavoc/id-go-client"
    "fmt"
)

var signupCommand = &Command {
        Name: "signup",
        Implementation: func(cmd *Command, args []string) {
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "Please provide this information for your new account",
                        Fields: []*prompt.Field {
                            {
                                Name: "fullname",
                                Title: "Fullname",
                                Instructions: "Enter your full name",
                                Shorthand: "fn",
                            },
                            {
                                Name: "email",
                                Title: "Email",
                                Instructions: "Enter your e-mail address",
                                Shorthand: "e",
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
            fullname, email, password, password_confirmation := form.Children["fullname"], form.Children["email"], form.Children["password"], form.Children["password_confirmation"]

            app := cmd.GetDeeqApplication()
            signUp := id.SignUp{
                Email: email.Value,
                Fullname: fullname.Value,
                Password: password.Value,
                PasswordConfirmation: password_confirmation.Value,
            }
            err := app.GetIdClient().SignUp(signUp)
            if err != nil {
                panic(err)
            }
            fmt.Printf(`
    Account successfully created.
    
    You will receive a new email soon with a confirmation code.

    Use the command 'confirm' to confirm your account with the confirmation code.

`)
            //cmd.GetApplication().LaunchCommand("confirm", 
        },
        Description: "Sign-up for a Free account",
        Help: `
    Sign-up for a Free account.

    Examples:
    
        $ deeq signup
        > Fullname: John Doe
        > Email: your_email@gmail.com
        > Password: your_password
        > Confirm Password: your_password

    or you could pass the values using arguments

        $ deeq signup --fullname --email=your_email@gmail.com --password=your_password --password_confirmation=your_password

    After the signup, you will receive an email with a confirmation code. Use the 'confirm' command with
    the provided code to use your account.

    See the 'deeq confirm' command for more information.
`,
    }
