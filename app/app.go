package app

var app *Application

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

func Get() *Application {
	if app == nil {
		app = new()
	}

	return app
}
