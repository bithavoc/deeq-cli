package main

import (
    prompt "github.com/bithavoc/goprompt"
    deeq "github.com/bithavoc/deeq-go-client"
    "fmt"
)

var createCommand = &Command {
        Name: "create",
        RequiresUser: true,
        Implementation: func(cmd *Command, args []string) {
            app := cmd.GetDeeqApplication()
            if app.GetCurrentUser().Token.Code == "" {
                return
            }
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "Enter the content of the task, include as many #hashtags as you need to organize it",
                        Fields: []*prompt.Field {
                            {
                                Name: "text",
                                Title: "Text",
                                Instructions: "Please enter the text of your new task and don't forget to use #one or #more hashtags",
                                Shorthand: "t",
                            },
                        },
                    },
                },
            }
            result := prompt.Process(args)
            form := result.Children["form.0"]
            textAnswer := form.Children["text"]

            tid := deeq.NewTaskId()
            stask, err := app.GetDeeqClient().SetTask(&deeq.Task {
                Id: tid,
                Text: textAnswer.Value,
                Status: deeq.TaskStatusIncomplete,
            })
            if err != nil {
                panic(err)
            }
            fmt.Println("Created ", stask.Id)
        },
        Description: "Creates a new task",
        Help: `

    You can use this command to create tasks.

    Examples:
    
        $ deeq create
        > Text: I have to create at least one task in #deeq with two or more #hashtags to be #happy

    Or, you can also do everything with a single command:

        $ deeq create --text "I have to create at least one task in #deeq with two ore more #hashtags to be #happy"

    Or, use the shorthand arguments:

        $ deeq create -t "I have to create at least one task in #deeq with two ore more #hashtags to be #happy"

    A reference or unique identifier for the task will be provided so you can use to complete or the delete the task.

    The hashtags you use in the text will be used as tags for your task.
    Tags are linked as parent-child and child-parent relationship.

    Example:

    Task Text: "#design #landing page for #AndroidApp"

    Tags will come out as:

        #design
        - #AndroidApp
        - #landing

        #landing
        - #AndroidApp
        - #design

        #AndroidApp
        - #landing
        - #design
     
    For more information on listing the tasks, see the command 'deeq all'.
`,
    }
