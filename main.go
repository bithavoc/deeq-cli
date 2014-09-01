package main

import (
    "fmt"
    "os"
    "io"
    "text/template"
    "strings"
)

type Command struct {
    Name string
    Description string
    Help string
    Implementation func(cmd *Command, args []string)
}

var commands = []*Command {
    loginCommand,
    {
        Name: "signup",
        Description: "Sign-up for a FREE Bithavoc account",
        Help: `signup a task a given task
            Examples:

        `,
    },
    {
        Name: "create",
        Description: "creates a new task",
        Help: `creates a task a given task
            Examples:
        `,
    },
    {
        Name: "delete",
        Description: "deletes a task",
        Help: `deletes a given task
            Examples:
        `,
    },
    {
        Name: "start",
        Description: "saves the task as current",
        Help: `starts a given task
            Examples:
        `,
    },
    {
        Name: "complete",
        Description: "completes a task (uses current task as default)",
        Help: `completes a given task

        `,
    },
    {
        Name: "list",
        Description: "list tasks in a specific hashtag",
        Help: `lists all the tasks in a specific hashtags

        `,
    },
    {
        Name: "sync",
        Description: "synchronize your local tasks with in the cloud",
        Help: `syncs the 

        local tasks against the cloud
        `,
    },
}

type AboutContent int16

const ShortAbout AboutContent = 0
const LongAbout AboutContent = 1

func printAbout(content AboutContent) {
    fmt.Println("Deeq 1.0.0")
    fmt.Println("2014 http://bithavoc.io")
    if(content == LongAbout) {
        fmt.Println("MIT License")
    }
}

var appArgs []string

func getAppArg(i int) string {
    if len(appArgs) < (i + 1) {
        return ""
    }
    return appArgs[i]
}

var usageTemplate = `Usage: deeq <command> [options] [arguments]

Commands:

    {{range .}}
        {{.Name | printf "%-11s"}} {{.Description}}{{end}}

Run 'deeq help [command]' for details.
`

func tmpl(w io.Writer, text string, data interface{}) {
    t := template.New("top")
    t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
    template.Must(t.Parse(text))
    if err:= t.Execute(w, data); err != nil {
        panic(err)
    }
}

func printUsage() {
    tmpl(os.Stderr, usageTemplate, commands)
}

func (cmd *Command)PrintUsage() {
    fmt.Fprintf(os.Stderr, cmd.Help)
}

func queryCommandByName(name string) *Command {
    for _, cmd := range commands {
        if cmd.Name == name {
            return cmd
        }
    }
    fmt.Fprintf(os.Stderr, "Unknown command %s\n", name)
    return nil
}

func main() {
    appArgs = os.Args[1:]
    mainCommandName := getAppArg(0)
    commandTarget := getAppArg(1)
    if mainCommandName == "" || mainCommandName == "help" {
        if commandTarget == "" {
            // print help
            printAbout(ShortAbout)
            printUsage()
        } else {
            // show help of specific command
            cmd := queryCommandByName(mainCommandName)
            if cmd == nil {
                printUsage()
            } else {
                // print command help
                cmd.PrintUsage()
            }
        }
        os.Exit(2)
    }
    cmd := queryCommandByName(mainCommandName)
    cmd.Implementation(cmd, appArgs[1:])
}
