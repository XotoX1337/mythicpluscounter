package main

import (
	"embed"
	"fmt"
	"github.com/XotoX1337/mythicpluscounter/app"
	"github.com/XotoX1337/mythicpluscounter/app/routes/index"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
	"net/http"
	"time"
)

//go:embed app/views
var views embed.FS

//go:embed public
var public embed.FS

func main() {

	App := app.Get()

	engine := html.NewFileSystem(http.FS(views), ".html")
	engine.AddFunc("Name", func() string {
		return App.Name
	})
	engine.AddFunc("Version", func() string {
		return App.Version
	})
	engine.AddFunc("Year", func() string {
		return fmt.Sprintf("%d", time.Now().Year())
	})

	fbr := fiber.New(fiber.Config{
		AppName: App.Name,
		Views:   engine,
	})

	fbr.Get("/", index.Index)

	fbr.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(public),
		PathPrefix: "public",
	}))

	log.Fatal(fbr.Listen(":5000"))
}
