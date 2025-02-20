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
	"chrisgriffis.com/model"

	"github.com/gofiber/fiber/v2"
)

type index struct{}

func (i index) RegisterRoutes(router fiber.Router) error {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Posts": []model.Linkable{
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
			"Projects": []model.Project{
				{
					Linkable: model.Linkable{
						Title:       "Project 1",
						Description: "This is a test of the emergency broadcast system.",
						Slug:        "project-1",
					},
					Technologies: []string{"Go", "Fiber", "HTML"},
				},
				{
					Linkable: model.Linkable{
						Title:       "Project 2",
						Description: "This is a test of the emergency broadcast system.",
						Slug:        "project-2",
					},
					Technologies: []string{"Go", "Fiber", "HTML"},
				},
				{
					Linkable: model.Linkable{
						Title:       "Project 3",
						Description: "This is a test of the emergency broadcast system.",
						Slug:        "project-3",
					},
					Technologies: []string{"Go", "Fiber", "HTML"},
				},
			},
		}, "layouts/main")
	})

	return nil
}

func NewIndex() RouteRegister {
	return &index{}
}
