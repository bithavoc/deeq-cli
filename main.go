package main

import (
    "fmt"
    "os"
    "io"
    "io/ioutil"
    "text/template"
    "strings"
    "errors"
    id "github.com/bithavoc/id-go-client"
    deeq "github.com/bithavoc/deeq-go-client"
    "path/filepath"
    "encoding/json"
    "math/rand"
    "time"
)

const appVersion = "1.0.0"

type BasicApplication struct {
    commands []*Command
    userLoggedIn bool
}

func (app *BasicApplication) IsUserLoggedIn() bool {
    return app.userLoggedIn
}

func (app *BasicApplication) SetUserIsLoggedIn(loggedIn bool) {
    app.userLoggedIn = loggedIn
}

type Application interface {
    PrintUsage()
    queryCommandByName(name string) *Command
    GetCommands() []*Command
    IsUserLoggedIn() bool
    SetUserIsLoggedIn(loggedIn bool)
}

func (app *BasicApplication) AddCommand(cmd *Command) {
    app.commands = append(app.commands, cmd)
}

func (app *BasicApplication) GetCommands() []*Command {
    return app.commands
}

func showNiceError(err interface{}) {
    fmt.Println("FAILED:", err)
}

func run(app Application) {
    defer func() {
        if r := recover(); r != nil {
            if _, ok := r.(*id.IdError); ok {
                showNiceError(r)
            } else if _, ok := r.(*deeq.DeeqError); ok {
                showNiceError(r)
            } else {
                panic(r)
            }
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
            app.PrintUsage()
        } else {
            // show help of specific command
            cmd := app.queryCommandByName(commandTarget)
            if cmd == nil {
                app.PrintUsage()
            } else {
                // print command help
                cmd.PrintUsage()
                fmt.Println("")
            }
        }
        os.Exit(2)
    }
    cmd := app.queryCommandByName(mainCommandName)
    if cmd == nil {
        app.PrintUsage()
        os.Exit(2)
    }
    if cmd.RequiresUser && !app.IsUserLoggedIn() {
        fmt.Fprintf(os.Stderr, `
    I'm sorry, please log-in first
    `)
        cmd := app.queryCommandByName("login")
        cmd.PrintUsage()
        os.Exit(2)
    }
    cmd.Implementation(cmd, appArgs[1:])
}

type Command struct {
    Name string
    Description string
    Help string
    Implementation func(cmd *Command, args []string)
    app Application
    RequiresUser bool
}

func (cmd *Command)GetApplication() Application {
    return cmd.app
}

type AboutContent int16

const ShortAbout AboutContent = 0
const LongAbout AboutContent = 1

func printAbout(content AboutContent) {
    fmt.Printf(`
    Deeq %s - 2014 bithavoc.io

    http://deeqapp.com

`, appVersion)
    if(content == LongAbout) {
        fmt.Println("MIT License")
    }
    fmt.Println("")
}

var appArgs []string

func getAppArg(i int) string {
    if len(appArgs) < (i + 1) {
        return ""
    }
    return appArgs[i]
}

var usageTemplate = `
    Usage: deeq <command> [options]

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
    fmt.Fprintf(os.Stderr, `
Command: %s
`, cmd.Name)
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
    SetCurrentUser(user id.User, save bool) error
    GetCurrentUser() id.User
    GetDeeqClient() *deeq.Client
}

type DeeqApp struct {
    BasicApplication
    Id id.Client
    CurrentUser id.User
    Deeq *deeq.Client
}

func (app * DeeqApp) GetIdClient() id.Client {
    return app.Id
}

func (app * DeeqApp) GetDeeqClient() *deeq.Client {
    return app.Deeq
}

func (cmd *Command) GetDeeqApplication() DeeqApplication {
    app := cmd.GetApplication().(DeeqApplication)
    return app
}

func (app *DeeqApp) GetCurrentUser() id.User {
    return app.CurrentUser
}

func (app *DeeqApp) SetCurrentUser(user id.User, save bool) error {
    app.CurrentUser = user
    if save {
        return app.SaveCurrentUser()
    }
    app.SetUserIsLoggedIn(app.CurrentUser.Token.Code != "")
    app.Deeq = deeq.NewClient(app.CurrentUser.Token)
    app.Deeq.ApplicationVersion = appVersion
    app.Deeq.ApplicationUpgradeChanged = func(c *deeq.Client) {
        if c.ApplicationUpgrade.Available {
            fmt.Printf("\n! This version of Deeq(v%s) is outdated\n", c.ApplicationUpgrade.Version)
            if c.ApplicationUpgrade.Message != "" {
                fmt.Printf("! %s\n", c.ApplicationUpgrade.Message)
                fmt.Printf("! Follow your upgrade procedure or visit DeeqApp.com for installation instructions\n\n")
            }
        }
    }
    return nil
}

func retrieveAppUserHome() string {
    home := os.Getenv("HOME")
    if home == "" {
        panic(errors.New("User doesn't have HOME? is this Windows? Da fuq is going on"))
    }
    return home
}

func getCurrentUserFilePath() string {
    return filepath.Join(retrieveAppUserHome(), ".deeq_user")
}

func (app *DeeqApp) SaveCurrentUser() error {
    user := app.GetCurrentUser()
    filePath := getCurrentUserFilePath()
    if user.Token.Code != "" {
        serialized, err := json.Marshal(user)
        if err != nil {
            return err
        }
        err = ioutil.WriteFile(filePath, serialized, os.ModePerm)
        return err
    } else {
        _, err := os.Stat(filePath)
        if err == nil {
            err := os.Remove(filePath)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func (app *DeeqApp) LoadCurrentUser() error {
    filePath := getCurrentUserFilePath()
    _, err := os.Stat(filePath)
    if err == nil {
        if fileContent, err := ioutil.ReadFile(filePath); err == nil {
            user := id.User{}
            if err := json.Unmarshal(fileContent, &user); err == nil {
                app.SetCurrentUser(user, false)
            } else {
                return err
            }
        } else {
            return err
        }
    }
    return nil
}

func main() {
    rand.Seed(time.Now().Unix())
    app := &DeeqApp {
        Id: id.NewClient("<app-id>"),
    }
    app.AddCommand(loginCommand)
    app.AddCommand(logoutCommand)
    app.AddCommand(createCommand)
    app.AddCommand(allCommand)
    app.AddCommand(completeCommand)
    app.AddCommand(deleteCommand)
    app.AddCommand(whoamiCommand)
    app.AddCommand(aboutCommand)
    app.AddCommand(signupCommand)
    if err:= app.LoadCurrentUser(); err != nil {
        panic(err)
    }
    run(app)
}

