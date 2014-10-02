package main

import (
    prompt "github.com/bithavoc/goprompt"
    "fmt"
)

var allCommand = &Command {
        Name: "all",
        RequiresUser: true,
        Implementation: func(cmd *Command, args []string) {
            app := cmd.GetDeeqApplication()
            if app.GetCurrentUser().Token.Code == "" {
                return
            }
            prompt := &prompt.Prompt {
                Forms : []*prompt.Form {
                    {
                        Title: "Enter the tag you want to list",
                        Fields: []*prompt.Field {
                            {
                                Name: "tag",
                                Title: "Tag Name",
                                Instructions: `Please enter the name of tag you want to list`,
                                Shorthand: "t",
                            },
                        },
                    },
                },
            }
            result := prompt.Process(args)
            form := result.Children["form.0"]
            tagAnswer := form.Children["tag"]

            rootTag := tagAnswer.Value
            tasks, err := app.GetDeeqClient().GetTasksInTags(rootTag, "")
            if err != nil {
                panic(err)
            }
            for _, task := range tasks {
                var st string
                if task.Status == 0 {
                    st = "Pending"
                } else {
                    st = "Done"
                }
                fmt.Printf("%s - %s (%s)\n", task.Id, task.Text, st)
            }
            if len(tasks) == 0 {
                fmt.Println("No tasks found :(")
            } else if len(tasks) == 1 {
                fmt.Println("1 task found")
            } else {
                fmt.Printf("%d tasks found\n", len(tasks))
            }
        },
        Description: "Lists all(pending and completed) tasks inside a tag",
        Help: `

    Use this command to retrieve all(pending and complete) tasks under a tag.

    Example, having a tag structure as follows:    

        #design
        - #AndroidApp
        - #landing

        #landing
        - #AndroidApp
        - #design

        #AndroidApp
        - #landing
        - #design

    You can use this command to retrieve tasks under a tag:

        $ deeq all #design

    It also works with sub-tags:

        $ deeq all design/landing
        
        $ deeq all AndroidApp/landing
        
        $ deeq all AndroidApp/design

    For more information on creating tasks, see the command 'deeq create'.
`,
    }
