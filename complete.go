package main

import (
    prompt "github.com/bithavoc/goprompt"
    deeq "github.com/bithavoc/deeq-go-client"
    "fmt"
)

var completeCommand = &Command {
        Name: "complete",
        Implementation: func(cmd *Command, args []string) {
            app := cmd.GetDeeqApplication()
            if app.GetCurrentUser().Token.Code == "" {
                return
            }
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "Enter the reference for the task you want to complete",
                        Fields: []*prompt.Field {
                            {
                                Name: "reference",
                                Title: "Task Reference",
                                Instructions: `Please enter the reference identifier for the task you want to complete, this identifier looks like 'GMWJGSAPGA' and you get it when the task is created`,
                            },
                        },
                    },
                },
            }
            result := prompt.Process()
            form := result.Children["form.0"]
            referenceAnswer := form.Children["reference"]

            referenceId := deeq.TaskId(referenceAnswer.Value)
            task, err := app.GetDeeqClient().GetTask(referenceId)
            if err != nil {
                panic(err)
            }
            task.Status = deeq.TaskStatusComplete
            stask, err := app.GetDeeqClient().SetTask(task)
            if err != nil {
                panic(err)
            }
            fmt.Println("Completed ", stask.Id)
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
