/*
Copyright © 2025 Chris Griffis <dev@chrisgriffis.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"chrisgriffis.com/internal/quotes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsFS embed.FS

type Linkable struct {
	Title       string
	Description string
	Slug        string
}

type Project struct {
	Linkable     Linkable
	Technologies []string
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	quoter := quotes.NewQuoteService()

	engine := html.NewFileSystem(http.FS(viewsFS), ".html")
	engine.AddFunc("Join", strings.Join)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Posts": []Linkable{
				{
					Title:       "Hello, World!",
					Description: "This is a test of the emergency broadcast system.",
					Slug:        "hello-world",
				},
				{
					Title:       "Goodbye, World!",
					Description: "This is a test of the emergency broadcast system.",
					Slug:        "goodbye-world",
				},
			},
			"Projects": []Project{
				{
					Linkable: Linkable{
						Title:       "Project 1",
						Description: "This is a test of the emergency broadcast system.",
						Slug:        "project-1",
					},
					Technologies: []string{"Go", "Fiber", "HTML"},
				},
				{
					Linkable: Linkable{
						Title:       "Project 2",
						Description: "This is a test of the emergency broadcast system.",
						Slug:        "project-2",
					},
					Technologies: []string{"Go", "Fiber", "HTML"},
				},
				{
					Linkable: Linkable{
						Title:       "Project 3",
						Description: "This is a test of the emergency broadcast system.",
						Slug:        "project-3",
					},
					Technologies: []string{"Go", "Fiber", "HTML"},
				},
			},
		}, "layouts/main")
	})

	app.Get("/whoami", func(c *fiber.Ctx) error {
		return c.Render("whoami", fiber.Map{}, "layouts/main")
	})

	app.Get("/quote", func(c *fiber.Ctx) error {
		quote, err := quoter.GetQuote()
		if err != nil {
			logger.Error("error getting quote", slog.Any("error", err))
			return c.Status(http.StatusInternalServerError).SendString("Error getting quote")
		}

		if len(quote) == 0 {
			logger.Error("no quotes returned")
			return c.Status(http.StatusInternalServerError).SendString("No quotes returned")
		}

		// Return the first quote
		return c.JSON(quote[0])
	})

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		logger.Error("Error starting server", slog.Any("error", err))
	}
}
