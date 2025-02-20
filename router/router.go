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

type RouteRegister interface {
	RegisterRoutes(router fiber.Router) error
}

type RouteRegisterFunc func(router fiber.Router) error

func (r RouteRegisterFunc) RegisterRoutes(router fiber.Router) error {
	return r(router)
}
