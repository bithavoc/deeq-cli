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
}

var commands = []*Command {
    {
        Name: "login",
        Description: "Log-in using your Bithavoc's credentials",
    },
    {
        Name: "signup",
        Description: "Sign-up for a FREE Bithavoc account",
    },
    {
        Name: "create",
        Description: "creates a new task",
    },
    {
        Name: "delete",
        Description: "deletes a task",
    },
    {
        Name: "start",
        Description: "saves the task as current",
    },
    {
        Name: "complete",
        Description: "completes a task (uses current task as default)",
    },
    {
        Name: "list",
        Description: "list task in a specific hashtag",
    },
    {
        Name: "sync",
        Description: "synchronize your local tasks with in the cloud",
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

func main() {
    appArgs = os.Args[1:]
    mainCommandName := getAppArg(0)
    if mainCommandName == "" {
        // print help
        printAbout(ShortAbout);
        printUsage();
        os.Exit(1)
    }
}
