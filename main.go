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

	"chrisgriffis.com/controller"
	"chrisgriffis.com/internal/quotes"
	"chrisgriffis.com/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var viewsFS embed.FS

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	engine := html.NewFileSystem(http.FS(viewsFS), ".html")
	engine.AddFunc("Join", strings.Join)

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./public")

	quoter := quotes.NewQuoteService()
	quoteCtrl := controller.MustQuote(controller.NewQuote(quoter))

	err := run(
		app,
		router.NewIndex(),
		router.NewWhoami(),
		router.RouteRegisterFunc(WithRoutePrefix("/api/v1")(router.NewQuotes(quoteCtrl))),
	)

	if err != nil {
		logger.Error("Error running app", slog.Any("error", err))
	}
}

func WithRoutePrefix(prefix string) func(f router.RouteRegister) func(r fiber.Router) error {
	return func(f router.RouteRegister) func(r fiber.Router) error {
		return func(r fiber.Router) error {
			r = r.Group(prefix)
			return f.RegisterRoutes(r)
		}
	}
}

func run(app *fiber.App, routers ...router.RouteRegister) error {
	// register provided routers
	for _, r := range routers {
		if err := r.RegisterRoutes(app); err != nil {
			return fmt.Errorf("[RegisterRoutes] %w", err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default to 8080
	}

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		return fmt.Errorf("[Listen] %w", err)
	}

	return nil
}
