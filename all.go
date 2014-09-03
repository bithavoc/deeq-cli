package main

import (
    prompt "github.com/bithavoc/goprompt"
    "fmt"
)

var allCommand = &Command {
        Name: "all",
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
                            },
                        },
                    },
                },
            }
            result := prompt.Process()
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
            fmt.Println("Tasks ", len(tasks))
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
