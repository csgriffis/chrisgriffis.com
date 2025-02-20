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

package router

import (
	"log/slog"
	"net/http"

	"chrisgriffis.com/controller"

	"github.com/gofiber/fiber/v2"
)

type quotes struct {
	controller *controller.Quote
}

func (q quotes) RegisterRoutes(router fiber.Router) error {
	router.Get("/quote", func(c *fiber.Ctx) error {
		quote, err := q.controller.GetQuote()
		if err != nil {
			slog.Error("error getting quote", slog.Any("error", err))
			return c.Status(http.StatusInternalServerError).SendString("Error getting quote")
		}

		// Return the first quote
		return c.JSON(quote)
	})

	return nil
}

func NewQuotes(c *controller.Quote) RouteRegister {
	return &quotes{controller: c}
}
