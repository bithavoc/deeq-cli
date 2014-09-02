package main

import (
    "fmt"
    "os"
    "io"
    "text/template"
    "strings"
    id "github.com/bithavoc/id-go-client"
)

type BasicApplication struct {
    commands []*Command
}

type Application interface {
    PrintUsage()
    queryCommandByName(name string) *Command
    GetCommands() []*Command
}

func (app *BasicApplication) AddCommand(cmd *Command) {
    app.commands = append(app.commands, cmd)
}

func (app *BasicApplication) GetCommands() []*Command {
    return app.commands
}

func run(app Application) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("FAILED:", r)
            os.Exit(1)
        } else {
            fmt.Println("OK")
        }
    }()
    if app.GetCommands() != nil {
        for _, cmd := range(app.GetCommands()) {
            cmd.app = app
        }
    }
    appArgs = os.Args[1:]
    mainCommandName := getAppArg(0)
    commandTarget := getAppArg(1)
    if mainCommandName == "" || mainCommandName == "help" {
        if commandTarget == "" {
            // print help
            printAbout(ShortAbout)
            app.PrintUsage()
        } else {
            // show help of specific command
            cmd := app.queryCommandByName(mainCommandName)
            if cmd == nil {
                app.PrintUsage()
            } else {
                // print command help
                cmd.PrintUsage()
            }
        }
        os.Exit(2)
    }
    cmd := app.queryCommandByName(mainCommandName)
    cmd.Implementation(cmd, appArgs[1:])
}

type Command struct {
    Name string
    Description string
    Help string
    Implementation func(cmd *Command, args []string)
    app Application
}

func (cmd *Command)GetApplication() Application {
    return cmd.app
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

func (app *BasicApplication) PrintUsage() {
    tmpl(os.Stderr, usageTemplate, app.GetCommands())
}

func (cmd *Command)PrintUsage() {
    fmt.Fprintf(os.Stderr, cmd.Help)
}

func (app *BasicApplication) queryCommandByName(name string) *Command {
    for _, cmd := range app.GetCommands() {
        if cmd.Name == name {
            return cmd
        }
    }
    fmt.Fprintf(os.Stderr, "Unknown command %s\n", name)
    return nil
}

type DeeqApplication interface {
    Application
    GetIdClient() id.Client
}

type DeeqApp struct {
    BasicApplication
    Id id.Client
}

func (app * DeeqApp) GetIdClient() id.Client {
    return app.Id
}

func main() {
    app := &DeeqApp {
        Id: id.NewClient("<app-id>"),
    }
    app.AddCommand(loginCommand)
    run(app)
}

/*


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
            },

*/
