package main

import (
    prompt "github.com/bithavoc/goprompt"
    deeq "github.com/bithavoc/deeq-go-client"
    "fmt"
)

var completeCommand = &Command {
        Name: "complete",
        RequiresUser: true,
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
                                Shorthand: "r",
                            },
                        },
                    },
                },
            }
            result := prompt.Process(args)
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
        Description: "Marks a task as completed",
        Help: `

    Use this command to mark tasks as completed.
    You only need the reference to the task you want to complete.

    Examples:
    
    $ deeq complete
    > Task Reference: blfotbknks

    Or you can also use a shorthand to provide the reference to the task to complete.

    $ deeq complete -r blfotbknks
`,
    }
