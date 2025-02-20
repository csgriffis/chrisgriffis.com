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

import "github.com/gofiber/fiber/v2"

type whoami struct{}

func (w *whoami) RegisterRoutes(router fiber.Router) error {
	router.Get("/whoami", func(c *fiber.Ctx) error {
		return c.Render("whoami", fiber.Map{}, "layouts/main")
	})

	return nil
}

func NewWhoami() RouteRegister {
	return &whoami{}
}
