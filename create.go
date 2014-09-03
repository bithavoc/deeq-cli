package main

import (
    prompt "github.com/bithavoc/goprompt"
    deeq "github.com/bithavoc/deeq-go-client"
    "fmt"
)

var createCommand = &Command {
        Name: "create",
        Implementation: func(cmd *Command, args []string) {
            app := cmd.GetDeeqApplication()
            if app.GetCurrentUser().Token.Code == "" {
                return
            }
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "New Task",
                        Fields: []*prompt.Field {
                            {
                                Name: "text",
                                Title: "Text",
                                Instructions: "Please enter the text of your new task and don't forget to use #one or #more hashtags",
                            },
                        },
                    },
                },
            }
            result := prompt.Process()
            form := result.Children["form.0"]
            textAnswer := form.Children["text"]

            tid := deeq.NewTaskId()
            err := app.GetDeeqClient().SetTask(tid, textAnswer.Value)
            if err != nil {
                panic(err)
            }
            fmt.Println("Created ", tid)
        },
        Description: "Creates a new task with the given #hashtags in the text",
        Help: `

        If you already logged in using deeq login, you can use this command to forget yourself

    Examples:
    
    $ deeq create
    > Text: I have to create at least one task in #deeq with two or more #hashtags to be #happy

    Or, you can also do everything with a single command:

    $ deeq create --text="I have to create at least one task in #deeq with two ore more #hashtags to be #happy"

`,
    }
