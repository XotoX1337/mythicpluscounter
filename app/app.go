package app

import (
	"fmt"
	"strings"
	"time"
)

var app *Application

const (
	INFO_LOG  = "INFO:"
	WARN_LOG  = "WARN:"
	ERROR_LOG = "ERROR:"
)

type Config struct {
}
type Application struct {
	Name        string
	Version     string
	Author      string
	Description string
	Config      Config
}

func new() *Application {
	return &Application{
		Name:        "Mythicpluscounter",
		Version:     "0.1.0",
		Author:      "Frederic Leist <https://github.com/XotoX1337>",
		Description: "See how many runs you did this week in m+",
	}
}

func (app *Application) UserAgent() string {
	return fmt.Sprintf("%s/%s", app.Name, app.Version)
}

func Get() *Application {
	if app == nil {
		app = new()
	}

	return app
}

func (app *Application) print(format string, a ...any) {

	msgParts := []string{
		fmt.Sprintf("[%s]", time.Now().Format(time.DateTime)),
		format,
		"\n",
	}

	fmt.Printf(strings.Join(msgParts, " "), a...)
}

func PrintInfo(format string, a ...any) {
	app.print(
		strings.Join([]string{INFO_LOG, format}, " "),
		a...,
	)
}
